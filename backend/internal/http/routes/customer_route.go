package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
)

func NewCustomerRoute(db *gorm.DB, prefix fiber.Router) {
	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	customerRoute := prefix.Group("/customers")
	customerRoute.Get("/", customerController.FindAll)
	customerRoute.Get("/:code", customerController.FindByCode)
	customerRoute.Post("/", customerController.Create)
	customerRoute.Patch("/:code", customerController.Update)
	customerRoute.Delete("/:code", customerController.Delete)
}
