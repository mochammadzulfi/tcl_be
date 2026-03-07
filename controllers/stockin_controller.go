package controllers

import (
	"net/http"
	"tcl_be/config"
	"tcl_be/models"
	"tcl_be/services"

	"github.com/gin-gonic/gin"
)

type StockInRequest struct {
	ProductID uint `json:"product_id"`
	Qty       int  `json:"qty"`
}

func GetStockIn(c *gin.Context) {

	var stockIns []models.StockIn
	config.DB.
		Preload("Items").
		Preload("Items.Product").
		Find(&stockIns)

	c.JSON(http.StatusOK, stockIns)
}

func CreateStockIn(c *gin.Context) {

	var req StockInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	stockIn := models.StockIn{
		Status: "CREATED",
	}
	config.DB.Create(&stockIn)
	item := models.StockInItem{
		StockInID: stockIn.ID,
		ProductID: req.ProductID,
		Qty:       req.Qty,
	}

	config.DB.Create(&item)

	c.JSON(http.StatusOK, stockIn)
}

func StartStockIn(c *gin.Context) {

	id := c.Param("id")
	var stockIn models.StockIn
	config.DB.First(&stockIn, id)
	stockIn.Status = "IN_PROGRESS"
	config.DB.Save(&stockIn)

	c.JSON(http.StatusOK, stockIn)
}

func CompleteStockIn(c *gin.Context) {

	id := c.Param("id")
	err := services.CompleteStockIn(id)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "Stock In Completed",
	})
}
