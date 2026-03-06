package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model

	ProductID uint
	Product   Product

	PhysicalStock int
	ReservedStock int
}
