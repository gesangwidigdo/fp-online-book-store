package router

import (
	"github.com/Djuanzz/go-template/controller"
	"github.com/gin-gonic/gin"
)

func Book(r *gin.Engine, bc controller.BookController) {
	routes := r.Group("/api/book")
	{
		routes.GET("/", bc.GetAll)
		routes.POST("/", bc.Create)
		routes.GET("/:slug", bc.GetBySlug)
		routes.PUT("/:id", bc.Update)
		routes.DELETE("/:id", bc.Delete)
	}
}
