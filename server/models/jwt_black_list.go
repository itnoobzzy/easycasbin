package models

import "akcasbin/server/global"

type JwtBlackList struct {
	global.BASE_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
