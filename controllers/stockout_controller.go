package controllers

import (
	"net/http"
	"tcl_be/config"
	"tcl_be/models"
	"tcl_be/services"

	"github.com/gin-gonic/gin"
)

type StockOutRequest struct {
	ProductID uint `json:"product_id"`
	Qty       int  `json:"qty"`
}

func GetStockOut(c *gin.Context) {
	
	var stockOuts []models.StockOut
	config.DB.
	Preload("Items").
	Preload("Items.Product").
	Find(&stockOuts)

	c.JSON(http.StatusOK, stockOuts)
}

func CreateStockOut(c *gin.Context) {

	var req StockOutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	tx := config.DB.Begin()
	var inventory models.Inventory
	tx.Where("product_id = ?", req.ProductID).First(&inventory)
	available := inventory.PhysicalStock - inventory.ReservedStock
	if available < req.Qty {
		c.JSON(http.StatusBadRequest, "Stock not enough")
		tx.Rollback()
		return
	}
	inventory.ReservedStock += req.Qty
	tx.Save(&inventory)
	stockOut := models.StockOut{
		Status: "DRAFT",
	}
	tx.Create(&stockOut)
	item := models.StockOutItem{
		StockOutID: stockOut.ID,
		ProductID:  req.ProductID,
		Qty:        req.Qty,
	}

	tx.Create(&item)
	tx.Commit()

	c.JSON(http.StatusOK, stockOut)
}

func ProcessStockOut(c *gin.Context) {

	id := c.Param("id")
	var stockOut models.StockOut
	config.DB.First(&stockOut, id)
	if stockOut.Status != "DRAFT" {
		c.JSON(http.StatusBadRequest, "Invalid status")
		return
	}
	stockOut.Status = "IN_PROGRESS"
	config.DB.Save(&stockOut)

	c.JSON(http.StatusOK, stockOut)
}

func CompleteStockOut(c *gin.Context) {

	id := c.Param("id")
	err := services.CompleteStockOut(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "Stock Out Completed",
	})
}

func CancelStockOut(c *gin.Context) {

	id := c.Param("id")
	tx := config.DB.Begin()
	var stockOut models.StockOut
	tx.First(&stockOut, id)

	if stockOut.Status != "DRAFT" && stockOut.Status != "IN_PROGRESS" {
		c.JSON(http.StatusBadRequest, "Cannot cancel")
		tx.Rollback()
		return
	}

	var items []models.StockOutItem
	tx.Where("stock_out_id = ?", id).Find(&items)
	for _, item := range items {

		var inventory models.Inventory

		tx.Where("product_id = ?", item.ProductID).First(&inventory)

		inventory.ReservedStock -= item.Qty

		tx.Save(&inventory)
	}

	stockOut.Status = "CANCELLED"
	tx.Save(&stockOut)
	tx.Commit()

	c.JSON(http.StatusOK, stockOut)
}
