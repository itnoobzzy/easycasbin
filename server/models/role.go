package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null;size:64;uniqueIndex:idx_domain_role"`
	Domain string `json:"domain" gorm:"size:127;uniqueIndex:idx_domain_role,unique"`
}

func (Role) TableName() string {
	return "role"
}
