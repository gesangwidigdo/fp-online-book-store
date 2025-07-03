package router

import (
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/middleware"
	"github.com/gin-gonic/gin"
)

func Payment(r *gin.Engine, pc controller.PaymentController) {
	routes := r.Group("/api/payment")
	{
		routes.POST("/", pc.Create)
		routes.POST("/standard", middleware.RequireAuth, pc.CreateStandard)
	}
}