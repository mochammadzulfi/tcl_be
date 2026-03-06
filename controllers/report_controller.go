package controllers

import (
	"net/http"
	"tcl_be/config"
	"tcl_be/models"

	"github.com/gin-gonic/gin"
)

func StockInReport(c *gin.Context) {

	var stockIns []models.StockIn

	config.DB.
		Preload("Items").
		Preload("Items.Product").
		Where("status = ?", "DONE").
		Find(&stockIns)

	c.JSON(http.StatusOK, stockIns)
}

func StockOutReport(c *gin.Context) {

	var stockOuts []models.StockOut

	config.DB.
		Preload("Items").
		Preload("Items.Product").
		Where("status = ?", "DONE").
		Find(&stockOuts)

	c.JSON(http.StatusOK, stockOuts)
}
