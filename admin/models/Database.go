package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInstance struct {
	Db *gorm.DB
}

var Database DBInstance

func DBConn() {
	var dsn string = "root:12345678@tcp(127.0.0.1:3306)/goweb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&Post{}, User{}, Category{})
	Database = DBInstance{Db: db}
}
