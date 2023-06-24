package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
	"inventory-management/backend/cmd/config"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
	"os"
)

func NewInitializedRoutes(configuration config.Config, logFile *os.File) (*fiber.App, error) {
	// Init app and middlewares
	app := fiber.New(middleware.FiberConfig())
	app.Use(etag.New())
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(middleware.XApiKeyMiddleware(configuration))
	app.Use(middleware.NewCORSMiddleware())
	app.Use(middleware.NewLoggerMiddleware(logFile))
	app.Use(middleware.NewCSRFMiddleware(configuration))

	// Init database
	db, err := config.NewPostgresSQLGorm(configuration)
	if err != nil {
		return nil, err
	}

	// Register the routes
	NewRoutes(db, app)

	return app, nil
}

func NewRoutes(db *gorm.DB, app *fiber.App) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)

	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	productQualityRepository := repository.NewProductQualityRepository(db)
	productRepository := repository.NewProductRepository(db)
	productQualityService := service.NewProductQualityService(productQualityRepository, productRepository)
	productQualityController := controller.NewProductQualityController(productQualityService)

	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	supplierRepository := repository.NewSupplierRepository(db)
	supplierService := service.NewSupplierService(supplierRepository)
	supplierController := controller.NewSupplierController(supplierService)

	app.Get("/", WelcomeHandler)

	prefix := app.Group("/api")
	prefix.Post("/login", authController.Login)
	prefix.Post("/register", authController.Register)

	app.Use(middleware.NewJWTMiddleware())

	customerRoute := prefix.Group("/customers")
	{
		customerRoute.Get("/", customerController.FindAll)
		customerRoute.Get("/:code", customerController.FindByCode)
		customerRoute.Post("/", customerController.Create)
		customerRoute.Patch("/:code", customerController.Update)
		customerRoute.Delete("/:code", customerController.Delete)
	}

	productQualities := prefix.Group("/product-qualities")
	{
		productQualities.Get("/", productQualityController.FindAll)
		productQualities.Get("/:code", productQualityController.FindAllByProductCode)
		productQualities.Delete("/:id", productQualityController.Delete)
	}

	products := prefix.Group("/products")
	{
		products.Get("/", productController.FindAll)
		products.Get("/:code", productController.FindByCode)
		products.Post("/", productController.Create)
		products.Patch("/:code", productController.Update)
		products.Delete("/:code", productController.Delete)
	}

	suppliers := prefix.Group("/suppliers")
	{
		suppliers.Get("/", supplierController.FindAll)
		suppliers.Get("/:code", supplierController.FindByCode)
		suppliers.Post("/", supplierController.Create)
		suppliers.Patch("/:code", supplierController.Update)
		suppliers.Delete("/:code", supplierController.Delete)
	}

	users := prefix.Group("/users")
	{
		users.Get("/", userController.FindAll)
		users.Get("/:id", userController.FindByID)
		users.Post("/", userController.Create)
		users.Patch("/:id", userController.Update)
		users.Delete("/:id", userController.Delete)
	}

	app.Get("*", NotFoundHandler)
}

func NotFoundHandler(c *fiber.Ctx) error {
	return response.ReturnJSON(c, fiber.StatusNotFound, "the requested resource was not found", nil)
}

func WelcomeHandler(c *fiber.Ctx) error {
	return response.ReturnJSON(c, fiber.StatusOK, "Welcome to the Inventory Management API", nil)
}
