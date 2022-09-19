package request

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	jwt.StandardClaims
}
