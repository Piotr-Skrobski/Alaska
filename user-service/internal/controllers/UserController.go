package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Piotr-Skrobski/Alaska/user-service/internal/dtos"
	"github.com/Piotr-Skrobski/Alaska/user-service/internal/services"
	"github.com/go-chi/chi/v5"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (uc *UserController) RegisterRoutes(r chi.Router) {
	r.Post("/users/register", uc.Register)
	r.Post("/users/login", uc.Login)

	r.Group(func(r chi.Router) {
		r.Use(uc.AuthMiddleware)
		r.Get("/users/me", uc.Me)
		r.Post("/users/logout", uc.Logout)
		r.Post("/users/delete", uc.Delete)
	})
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req dtos.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := uc.UserService.Register(r.Context(), req); err != nil {
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

	resp, err := uc.UserService.Login(r.Context(), req)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	uc.UserService.SessionService.SetSessionCookie(w, resp.Token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (uc *UserController) Me(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	session, err := uc.UserService.GetSession(r.Context(), cookie.Value)
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

	uc.UserService.Logout(r.Context(), cookie.Value)
	uc.UserService.SessionService.ClearSessionCookie(w)

	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	session, err := uc.UserService.GetSession(r.Context(), cookie.Value)
	if err != nil {
		http.Error(w, "invalid session", http.StatusUnauthorized)
		return
	}

	userID, parseErr := strconv.Atoi(session.UserID)
	if parseErr != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	err = uc.UserService.Delete(r.Context(), userID)
	if err != nil {
		http.Error(w, "failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	uc.UserService.SessionService.ClearSessionCookie(w)
	w.WriteHeader(http.StatusNoContent)
}

func (uc *UserController) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		_, err = uc.UserService.GetSession(r.Context(), cookie.Value)
		if err != nil {
			http.Error(w, "invalid session", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
