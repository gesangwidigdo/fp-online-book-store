package main

import (
	"fmt"
	"os"

	"github.com/Djuanzz/go-template/config"
	"github.com/Djuanzz/go-template/controller"
	"github.com/Djuanzz/go-template/middleware"
	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/repository"
	"github.com/Djuanzz/go-template/router"
	"github.com/Djuanzz/go-template/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDatabase()
	client *snap.Client = config.ConnectMidtrans()

	userRepo repository.UserRepository = repository.NewUserRepository(db)
	paymentRepo repository.PaymentRepository = repository.NewPaymentRepository(db)

	userService service.UserService = service.NewUserService(userRepo)
	midtransService service.MidtransService = service.NewMidtransService(client)
	paymentService service.PaymentService = service.NewPaymentService(paymentRepo)

	userController controller.UserController = controller.NewUserController(userService)
	paymentController controller.PaymentController = controller.NewPaymentController(paymentService, midtransService)
)

func main() {
	fmt.Println("Final Project PBKK D")
	fmt.Println("Developed by:")
	fmt.Println("Adnan Abdullah Juan | 5025221155")
	fmt.Println("Muhammad Gesang Ridho Widigdo | 5025221216")

	defer config.CloseDatabase(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	router.User(server, userController)
	router.Payment(server, paymentController)

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
