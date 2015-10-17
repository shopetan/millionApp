package models

import (
  "fmt"
  "time"
_ "github.com/go-sql-driver/mysql" // 新規追加
  "github.com/jinzhu/gorm" // 新規追加
  "strings" // 新規追加
  "fmt" // 新規追加
)

type Item struct {
  itemID            string    `json:"itemID"`
  itemSupplier      string    `json:"itemSupplier"`
  itemStockQuantity string    `json:"itemStockQuantity"`
  itemBasePrice     string    `json:"itemBasePrice"`
  itemTags          string    `json:"itemTags"`
}