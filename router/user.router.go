package router

import (
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, uc controller.UserController) {
	routes := r.Group("/api/user")
	{
		routes.GET("/", uc.GetAll)
		routes.POST("/register", uc.Register)
		routes.POST("/login", uc.Login)
		routes.GET("/me", middleware.RequireAuth("user, admin"), uc.Me)
	}
}
