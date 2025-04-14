package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/dtos"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/models"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/repositories"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserRepo       *repositories.UserRepository
	SessionService *services.SessionService
}

func NewUserController(userRepo *repositories.UserRepository, sessionService *services.SessionService) *UserController {
	return &UserController{
		UserRepo:       userRepo,
		SessionService: sessionService,
	}
}

func (uc *UserController) RegisterRoutes(r chi.Router) {
	r.Post("/users/register", uc.Register)
	r.Post("/users/login", uc.Login)

	r.Group(func(r chi.Router) {
		r.Use(uc.AuthMiddleware)
		r.Get("/users/me", uc.Me)
		r.Post("/users/logout", uc.Logout)
	})
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}

	if err := uc.UserRepo.CreateUser(user); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var req dtos.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := uc.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if !uc.UserRepo.VerifyPassword(user, req.Password) {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := uc.SessionService.CreateSession(context.Background(), user)
	if err != nil {
		http.Error(w, "failed to create session", http.StatusInternalServerError)
		return
	}

	uc.SessionService.SetSessionCookie(w, token)

	resp := dtos.AuthResponse{
		Token: token,
		User:  *user,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	session, err := uc.SessionService.GetSession(context.Background(), cookie.Value)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func (uc *UserController) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uc.SessionService.DeleteSession(context.Background(), cookie.Value)
	uc.SessionService.ClearSessionCookie(w)

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		_, err = uc.SessionService.GetSession(context.Background(), cookie.Value)
		if err != nil {
			http.Error(w, "invalid session", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
