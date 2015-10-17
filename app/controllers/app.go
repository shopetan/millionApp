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
    for i := 0; i < 10; i++ {
        DbMap.Insert(&models.User{0, fmt.Sprintf("user%d", i)})
    }

    rows, _ := DbMap.Select(models.User{}, "select * from user")
    for _, row := range rows {
        user := row.(*models.User)
        fmt.Printf("%d, %s\n", user.Id, user.Name)
    }

    return c.Render()
}