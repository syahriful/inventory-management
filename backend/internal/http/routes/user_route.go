package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
)

func NewUserRoutes(db *gorm.DB, prefix fiber.Router) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	users := prefix.Group("/users")
	users.Get("/", userController.FindAll)
	users.Get("/:id", userController.FindByID)
	users.Post("/", userController.Create)
	users.Patch("/:id", userController.Update)
	users.Delete("/:id", userController.Delete)
}
