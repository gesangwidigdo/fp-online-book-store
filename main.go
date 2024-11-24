package main

import (
	"fmt"
	"os"

	"github.com/Djuanzz/go-template/config"
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/repository"
	"github.com/Djuanzz/go-template/router"
	"github.com/Djuanzz/go-template/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDatabase()

	userRepo repository.UserRepository = repository.NewUserRepository(db)

	userService service.UserService = service.NewUserService(userRepo)

	userController controller.UserController = controller.NewUserController(userService)
)

func main() {
	fmt.Println("Personal Template for Go project")

	defer config.CloseDatabase(db)

	server := gin.Default()
	router.User(server, userController)

	if err := model.Migrate(db); err != nil {
		panic("Failed to migrate database")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := server.Run(":" + port); err != nil {
		panic(err.Error())
	}

}
