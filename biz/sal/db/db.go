package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:123456@tcp(127.0.0.1:3306)/ddd_parking?charset=utf8mb4&parseTime=True&loc=Local"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
