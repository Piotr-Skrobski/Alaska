package dtos

import "github.com/Piotr-Skrobski/Alaska/user-service/internal/models"

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}
