package models

import (
_ "github.com/go-sql-driver/mysql" // 新規追加
  "github.com/jinzhu/gorm" // 新規追加
  "strings" // 新規追加
  "fmt" // 新規追加
)

type User struct {
  userID           string    `json:"userID"`
  userCompany      string    `json:"userCompany"`
  userDiscountRate int32    `json:"userDiscountRate"`
  GuestStaringRole string    `json:"guest_staring_role"`
}