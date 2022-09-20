package models

import "akcasbin/global"

type JwtBlackList struct {
	global.BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
