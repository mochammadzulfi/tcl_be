package models

import "gorm.io/gorm"

type StockIn struct {
	gorm.Model

	Status string

	Items []StockInItem
}

type StockInItem struct {
	gorm.Model

	StockInID uint
	ProductID uint

	Product Product

	Qty int
}
