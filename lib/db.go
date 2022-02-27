package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectionDB() *gorm.DB {
	dbCredentials := GetEnv("USER") + ":" + GetEnv("PASSWORD") + "@tcp(" + GetEnv("HOST") + ":" + GetEnv("PORT") + ")/" + GetEnv("DATABASE") + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dbCredentials), &gorm.Config{})
	if err != nil {
		return nil
	}
	return db
}

func DBConnection() *gorm.DB {
	return DBConn
}
