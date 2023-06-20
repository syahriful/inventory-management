package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
)

func NewAuthRoute(db *gorm.DB, prefix fiber.Router) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authController := controller.NewAuthController(userService)

	prefix.Post("/login", authController.Login)
	prefix.Post("/register", authController.Register)
}
