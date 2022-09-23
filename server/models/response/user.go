package response

import "akcasbin/models"

type LoginResponse struct {
	User      models.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}

type RegisterResponse struct {
	User models.User `json:"user"`
}
