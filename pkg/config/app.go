package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	d, err := gorm.Open("")
	if err != nil {
		panic(err)
	}
	db = d

}

func GetDb() *gorm.DB {
	return db
}
