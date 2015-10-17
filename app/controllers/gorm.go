// app/controllers/gorm.go
package controllers

import (
        _ "github.com/go-sql-driver/mysql"
        "github.com/jinzhu/gorm"
        "github.com/revel/revel"
        "github.com/shopetan/millionApp/app/models"
        "log"
)

var DB *gorm.DB

func InitDB() {
        db, err := gorm.Open("mysql", dbInfoString())
        if err != nil {
                log.Panicf("Failed to connect to database: %v\n", err)
        }
        db.DB()
	db.AutoMigrate(&models.user{})
	db.AutoMigrate(&models.item{})
	db.AutoMigrate(&models.order{})
        DB = &db
}

func dbInfoString() string {
        s, b := revel.Config.String("db.info")
        if !b {
                log.Panicf("database info not found")
        }

        return s
}