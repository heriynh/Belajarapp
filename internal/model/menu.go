package model

import "github.com/rocksus/go-restaurant-app/internal/model/constant"

type MenuItem struct {
	Nama      string
	OrderCode string
	Price     int64
	Type      constant.MenuType
}
