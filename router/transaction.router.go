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
	}
}