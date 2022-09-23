package utils

import (
	"akcasbin/global"
	"akcasbin/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.CASBIN_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims models.BaseClaims) models.CustomClaims {
	claims := models.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.CASBIN_CONFIG.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.CASBIN_CONFIG.JWT.ExpiresTime,
			Issuer:    global.CASBIN_CONFIG.JWT.Issuer,
			NotBefore: time.Now().Unix() - 1000,
		},
	}
	return claims
}

// CreateToken 创建JWT token
func (j *JWT) CreateToken(claims models.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 生成新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims models.CustomClaims) (string, error) {
	v, err, _ := global.CASBIN_CONCURRENCY_CONTROL.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析token
func (j *JWT) ParseToken(tokenString string) (*models.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	}
	return nil, TokenInvalid
}
