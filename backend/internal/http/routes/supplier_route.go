package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
)

func NewSupplierRoute(db *gorm.DB, prefix fiber.Router) {
	supplierRepository := repository.NewSupplierRepository(db)
	supplierService := service.NewSupplierService(supplierRepository)
	supplierController := controller.NewSupplierController(supplierService)

	suppliers := prefix.Group("/suppliers")
	suppliers.Get("/", supplierController.FindAll)
	suppliers.Get("/:code", supplierController.FindByCode)
	suppliers.Post("/", supplierController.Create)
	suppliers.Patch("/:code", supplierController.Update)
	suppliers.Delete("/:code", supplierController.Delete)
}
