package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	SKU      string `gorm:"unique"`
	Name     string
	Customer string
}
