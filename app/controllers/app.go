package controllers

import (
    "github.com/revel/revel"
    _ "github.com/go-sql-driver/mysql" // 新規追加
    _ "github.com/jinzhu/gorm" // 新規追加
    _ "strings" // 新規追加
    _ "fmt" // 新規追加
)

type App struct {
	*revel.Controller
}


func (c App) Index() revel.Result {
    return c.Render()
}
