package config

import (
	"github.com/MichaelSitanggang/bookstore/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/bookstore"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(entities.User{})
	return db
}
