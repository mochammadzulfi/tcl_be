package services

import (
	"errors"
	"tcl_be/config"
	"tcl_be/models"
)

func CompleteStockIn(id string) error {

	tx := config.DB.Begin()

	var stockIn models.StockIn
	tx.First(&stockIn, id)

	if stockIn.Status != "IN_PROGRESS" {
		tx.Rollback()
		return errors.New("invalid status")
	}

	var items []models.StockInItem
	tx.Where("stock_in_id = ?", id).Find(&items)

	for _, item := range items {

		var inventory models.Inventory

		result := tx.Where("product_id = ?", item.ProductID).First(&inventory)

		beforeStock := 0

		if result.Error != nil {

			inventory = models.Inventory{
				ProductID:     item.ProductID,
				PhysicalStock: item.Qty,
			}

			tx.Create(&inventory)

		} else {

			beforeStock = inventory.PhysicalStock

			inventory.PhysicalStock += item.Qty
			tx.Save(&inventory)
		}

		log := models.StockLog{
			ProductID:       item.ProductID,
			TransactionType: "STOCK_IN",
			ReferenceID:     stockIn.ID,
			Qty:             item.Qty,
			BeforeStock:     beforeStock,
			AfterStock:      beforeStock + item.Qty,
		}

		tx.Create(&log)
	}

	stockIn.Status = "DONE"

	tx.Save(&stockIn)

	tx.Commit()

	return nil
}
