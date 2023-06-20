package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
)

func NewProductQualityRoute(db *gorm.DB, prefix fiber.Router) {
	productQualityRepository := repository.NewProductQualityRepository(db)
	productRepository := repository.NewProductRepository(db)
	productQualityService := service.NewProductQualityService(productQualityRepository, productRepository)
	productQualityController := controller.NewProductQualityController(productQualityService)

	productQualities := prefix.Group("/product-qualities")
	productQualities.Get("/", productQualityController.FindAll)
	productQualities.Get("/:code", productQualityController.FindAllByProductCode)
	productQualities.Delete("/:id", productQualityController.Delete)
}
