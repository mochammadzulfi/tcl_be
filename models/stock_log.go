package models

import "gorm.io/gorm"

type StockLog struct {
	gorm.Model

	ProductID uint

	TransactionType string

	ReferenceID uint

	Qty int

	BeforeStock int
	AfterStock  int
}
