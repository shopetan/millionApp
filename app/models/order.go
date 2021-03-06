package models

import (
_ "github.com/go-sql-driver/mysql" // 新規追加
_  "github.com/jinzhu/gorm" // 新規追加
  "strings" // 新規追加
)

type Item struct {
  orderID           string    `json:"orderID"`
  orderDataTime     string    `json:"orderDataTime"`
  orderUserID       string    `json:"orderUserID"`
  orderItemID       string    `json:"orderItemID"`
  orderQuantity     string    `json:"orderQuantity"`
  orderState        string    `json:"orderState"`
  orderTags         string    `json:"orderTags"`
}