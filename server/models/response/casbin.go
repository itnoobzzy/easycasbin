package response

import "akcasbin/models"

type AddRoleForUserInDomain struct {
	User      models.User `json:"user"`
	Token     string      `json:"token"`
	ExpiresAt int64       `json:"expiresAt"`
}
