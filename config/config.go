package config

import (
	"github.com/anovanmaximuz/go-jwt/structs"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:@tcp(128.199.211.144:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql","root:@(localhsot:3306)/popfren?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
