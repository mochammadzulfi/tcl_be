package controllers

import (
	"net/http"
	"tcl_be/config"
	"tcl_be/models"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {

	var inventories []models.Inventory
	config.DB.Preload("Product").Find(&inventories)
	var result []gin.H

	for _, inv := range inventories {
		available := inv.PhysicalStock - inv.ReservedStock
		result = append(result, gin.H{
			"product_id":      inv.ProductID,
			"physical_stock":  inv.PhysicalStock,
			"reserved_stock":  inv.ReservedStock,
			"available_stock": available,
		})
	}

	c.JSON(http.StatusOK, result)
}
