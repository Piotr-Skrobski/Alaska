package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/models"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type SessionService struct {
	redisClient *redis.Client
	sessionTTL  time.Duration
}

func NewSessionService(redisClient *redis.Client) *SessionService {
	return &SessionService{
		redisClient: redisClient,
		sessionTTL:  6 * time.Hour,
	}
}

func (s *SessionService) GenerateSessionToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "according to rand.Read, it should never return an error - something interesting happened", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *SessionService) CreateSession(ctx context.Context, user *models.User) (string, error) {
	token, err := s.GenerateSessionToken()
	if err != nil {
		return "", fmt.Errorf("failed to generate session token: %w", err)
	}

	sessionData := map[string]interface{}{
		"user_id":    user.ID,
		"email":      user.Email,
		"username":   user.Username,
		"created_at": time.Now().Unix(),
	}

	err = s.redisClient.HSet(ctx, "session:"+token, sessionData).Err()
	if err != nil {
		return "", fmt.Errorf("failed to store session: %w", err)
	}

	err = s.redisClient.Expire(ctx, "session:"+token, s.sessionTTL).Err()
	if err != nil {
		return "", fmt.Errorf("failed to set session expiration: %w", err)
	}

	return token, nil
}

func (s *SessionService) GetSession(ctx context.Context, token string) (map[string]string, error) {
	sessionData, err := s.redisClient.HGetAll(ctx, "session:"+token).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	if len(sessionData) == 0 {
		return nil, fmt.Errorf("session not found")
	}

	return sessionData, nil
}

func (s *SessionService) DeleteSession(ctx context.Context, token string) error {
	return s.redisClient.Del(ctx, "session:"+token).Err()
}

func (s *SessionService) ExtendSession(ctx context.Context, token string) error {
	exists, err := s.redisClient.Exists(ctx, "session:"+token).Result()
	if err != nil {
		return err
	}
	if exists == 0 {
		return fmt.Errorf("session not found")
	}

	return s.redisClient.Expire(ctx, "session:"+token, s.sessionTTL).Err()
}

func (s *SessionService) SetSessionCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(s.sessionTTL.Seconds()),
	})
}

func (s *SessionService) ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})
}
