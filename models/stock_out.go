package models

import "gorm.io/gorm"

type StockOut struct {
	gorm.Model

	Status string

	Items []StockOutItem
}

type StockOutItem struct {
	gorm.Model

	StockOutID uint
	ProductID  uint

	Product Product

	Qty int
}
