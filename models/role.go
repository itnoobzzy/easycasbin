package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null;size:64;uniqueIndex:domain_role"`
	Domain string `json:"domain" gorm:"size:127;uniqueIndex:domain_role"`
}

func (Role) TableName() string {
	return "role"
}
