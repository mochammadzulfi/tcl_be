package routes

import (
	"tcl_be/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	//product
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)

	//inventory
	r.GET("/inventory", controllers.GetInventory)

	//stock in
	r.GET("/stock-in", controllers.GetStockIn)
	r.POST("/stock-in", controllers.CreateStockIn)
	r.PATCH("/stock-in/:id/start", controllers.StartStockIn)
	r.PATCH("/stock-in/:id/complete", controllers.CompleteStockIn)

	//stock out
	r.GET("/stock-out", controllers.GetStockOut)
	r.POST("/stock-out", controllers.CreateStockOut)
	r.PATCH("/stock-out/:id/process", controllers.ProcessStockOut)
	r.PATCH("/stock-out/:id/complete", controllers.CompleteStockOut)
	r.PATCH("/stock-out/:id/cancel", controllers.CancelStockOut)

	//report
	r.GET("/reports/stock-in", controllers.StockInReport)
	r.GET("/reports/stock-out", controllers.StockOutReport)

}
