package helper

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "3306"
	user     = "root"
	password = "password"
	dbname   = "user_manager"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewGormDB(dbConfig DBConfig) *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
