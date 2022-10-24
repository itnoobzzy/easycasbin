package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null;size:64"`
	Domain string `json:"domain" gorm:"size:127"`
}

func (Role) TableName() string {
	return "role"
}
