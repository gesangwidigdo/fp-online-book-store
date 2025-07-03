package router

import (
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/middleware"
	"github.com/gin-gonic/gin"
)

func Transaction(r *gin.Engine, tc controller.TransactionController) {
	routes := r.Group("/api/transaction")
	{
		routes.POST("/", middleware.RequireAuth, tc.Create)
		routes.GET("/status", middleware.RequireAuth, tc.GetTransactionStatus)
		routes.GET("/list", middleware.RequireAuth,tc.GetAllTransactionByUserLogin)
		routes.GET("/:id", tc.GetTransactionWithBooksByID)
		routes.PUT("/total/:id", tc.CalculateGrandTotal)
		routes.GET("/", middleware.RequireAuth, tc.GetTransactionWithBooksByUserLogin)
		routes.PUT("/total", middleware.RequireAuth, tc.CalculateGrandTotalByUserLogin)
	}
}