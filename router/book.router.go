package router

import (
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/middleware"
	"github.com/gin-gonic/gin"
)

func Book(r *gin.Engine, bc controller.BookController) {
	routes := r.Group("/api/book")
	{
		routes.GET("/", bc.GetAll)
		routes.POST("/", middleware.RequireAuth("admin"), bc.Create)
		routes.GET("/:slug", bc.GetBySlug)
		routes.PUT("/:id", middleware.RequireAuth("admin"), bc.Update)
		routes.DELETE("/:id", middleware.RequireAuth("admin"), bc.Delete)
	}
}
