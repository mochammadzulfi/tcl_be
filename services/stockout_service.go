package services

import (
	"errors"
	"tcl_be/config"
	"tcl_be/models"
)

func CompleteStockOut(id string) error {

	tx := config.DB.Begin()

	var stockOut models.StockOut
	tx.First(&stockOut, id)

	if stockOut.Status != "IN_PROGRESS" {
		tx.Rollback()
		return errors.New("invalid status")
	}

	var items []models.StockOutItem
	tx.Where("stock_out_id = ?", id).Find(&items)

	for _, item := range items {

		var inventory models.Inventory
		tx.Where("product_id = ?", item.ProductID).First(&inventory)

		beforeStock := inventory.PhysicalStock

		inventory.PhysicalStock -= item.Qty
		inventory.ReservedStock -= item.Qty

		tx.Save(&inventory)

		log := models.StockLog{
			ProductID:       item.ProductID,
			TransactionType: "STOCK_OUT",
			ReferenceID:     stockOut.ID,
			Qty:             item.Qty,
			BeforeStock:     beforeStock,
			AfterStock:      beforeStock - item.Qty,
		}

		tx.Create(&log)
	}

	stockOut.Status = "DONE"

	tx.Save(&stockOut)

	tx.Commit()

	return nil
}
