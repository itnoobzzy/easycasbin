package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
}
