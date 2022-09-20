package models

import (
	"akcasbin/global"
)

type CasbinRule struct {
	global.BASE_MODEL
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `json:"ptype" gorm:"size:1"`
	V0    string `json:"v0" gorm:"size:127"`
	V1    string `json:"v1" gorm:"size:127"`
	V2    string `json:"v2" gorm:"size:127"`
	V3    string `json:"v3" gorm:"size:127"`
	V4    string `json:"v4" gorm:"size:127"`
	V5    string `json:"v5" gorm:"size:127"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
