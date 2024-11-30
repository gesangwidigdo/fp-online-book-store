package router

import (
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/middleware"
	"github.com/gin-gonic/gin"
)

func BookTransaction(r *gin.Engine, btc controller.BookTransactionController) {
	routes := r.Group("/api/book_transaction")
	{
		routes.POST("/", middleware.RequireAuth("user"), btc.Create)
		routes.GET("/:transaction_id", middleware.RequireAuth("user"), btc.GetByTransactionID)
		routes.PUT("/:transaction_id", middleware.RequireAuth("user"), btc.UpdateQuantity)
		routes.DELETE("/:transaction_id", middleware.RequireAuth("user"), btc.Delete)
	}
}
