package controllers

import (
	"net/http"
	"tcl_be/config"
	"tcl_be/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

func GetProducts(c *gin.Context) {

	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}
