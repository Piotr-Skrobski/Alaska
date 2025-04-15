package services

import (
	"context"
	"errors"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/dtos"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/models"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/repositories"
)

type UserService struct {
	UserRepo       *repositories.UserRepository
	SessionService *SessionService
}

func NewUserService(userRepo *repositories.UserRepository, sessionService *SessionService) *UserService {
	return &UserService{
		UserRepo:       userRepo,
		SessionService: sessionService,
	}
}

func (us *UserService) Register(ctx context.Context, req dtos.RegisterRequest) error {
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}
	return us.UserRepo.CreateUser(user)
}

func (us *UserService) Login(ctx context.Context, req dtos.LoginRequest) (*dtos.AuthResponse, error) {
	user, err := us.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !us.UserRepo.VerifyPassword(user, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := us.SessionService.CreateSession(ctx, user)
	if err != nil {
		return nil, err
	}

	return &dtos.AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (us *UserService) GetSession(ctx context.Context, token string) (*dtos.Session, error) {
	return us.SessionService.GetSession(ctx, token)
}

func (us *UserService) Logout(ctx context.Context, token string) {
	us.SessionService.DeleteSession(ctx, token)
}
