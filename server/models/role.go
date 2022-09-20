package models

import (
	"akcasbin/global"
)

type Role struct {
	global.BASE_MODEL
	Name   string `json:"name" gorm:"not null;size:64"`
	Domain string `json:"domain" gorm:"size:127"`
}

func (Role) TableName() string {
	return "role"
}
