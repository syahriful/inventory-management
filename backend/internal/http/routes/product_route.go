package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
)

func NewProductRoute(db *gorm.DB, prefix fiber.Router) {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	products := prefix.Group("/products")
	products.Get("/", productController.FindAll)
	products.Get("/:id", productController.FindByID)
	products.Post("/", productController.Create)
	products.Patch("/:id", productController.Update)
	products.Delete("/:id", productController.Delete)
}
