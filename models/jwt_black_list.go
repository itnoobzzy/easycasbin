package models

import (
	"gorm.io/gorm"
)

type JwtBlackList struct {
	gorm.Model
	Jwt string `gorm:"type:text;comment:jwt"`
}
