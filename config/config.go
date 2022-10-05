package config

import (
	"fmt"
	"practice/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Dragoncit1234@tcp(127.0.0.1:3306)/latihan?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(fmt.Sprintf("failed to connect to db: %+v", err))
	}

	db.AutoMigrate(structs.Person{})
	return db
}
