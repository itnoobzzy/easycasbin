package models

import (
	"database/sql/driver"
	"easycasbin/global"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type User struct {
	gorm.Model
	FirstName      string      `json:"first_name" gorm:"type:text"`
	LastName       string      `json:"last_name" gorm:"type:text"`
	Username       string      `json:"username" gorm:"type:text"`
	Password       string      `json:"password" gorm:"comment:密码"`
	Active         int         `json:"active" gorm:"type:tinyint(1);not null;comment:是否激活"`
	Email          string      `json:"email" gorm:"type:text;comment:邮箱"`
	LastLogin      *LocalTime  `json:"last_login" gorm:"type:time;comment:last_login"`
	LoginCount     int         `json:"login_count" gorm:"size:11;comment:login_count"`
	FailLoginCount int         `json:"fail_login_count" gorm:"size:11;comment:fail_login_count"`
	Params         global.JSON `json:"params" gorm:"type:json;serializer:json"`
	CreatedByFk    int         `json:"created_by_fk" gorm:"size:11;comment:创建人id"`
	ChangedByFk    int         `json:"changed_by_fk" gorm:"size:11;comment:修改人id"`
	RegisterFrom   string      `json:"register_from" gorm:"type:text"`
	AliasName      string      `json:"alias_name" gorm:"size:64"`
	DepartmentPath string      `json:"department_path" gorm:"size:200"`
	Mobile         string      `json:"mobile" gorm:"size:20"`
	Position       string      `json:"position" gorm:"size:100"`
	ThumbAvatar    string      `json:"thumb_avatar" gorm:"size:1000"`
	WxUsername     string      `json:"wx_username" gorm:"size:100"`
}

func (User) TableName() string {
	return "user"
}
