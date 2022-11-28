package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User2 struct {
	gorm.Model
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}

func main() {
	db, err := gorm.Open("mysql", "dbadmin:hE4sqSfuCQeXEXwz@(rm-3nsc58907o3epw2me.mysql.rds.aliyuncs.com:3306)/easycasbin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&User2{})

	user := User2{Name: "Jinzhu", Age: 18}
	db.Create(&user)

	var u User2
	db.Where("id=?", 1).Find(&u)
	fmt.Println(&user)

	db.Delete(&User2{}, 1)
}
