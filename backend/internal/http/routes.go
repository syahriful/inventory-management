package http

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/gorm"
	"inventory-management/backend/cmd/config"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/response"
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

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}

	// Register the routes
	NewRoutes(db, app, es)

	return app, nil
}

func NewRoutes(db *gorm.DB, app *fiber.App, es *elasticsearch.Client) {
	// Init repositories
	userRepository := repository.NewUserRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	productQualityRepository := repository.NewProductQualityRepository(db)
	productRepository := repository.NewProductRepository(db)
	supplierRepository := repository.NewSupplierRepository(db)
	txRepository := repository.NewTxRepository(db, transactionRepository, productQualityRepository)

	// Init services
	userService := service.NewUserService(userRepository)
	customerService := service.NewCustomerService(customerRepository)
	productQualityService := service.NewProductQualityService(productQualityRepository, productRepository)
	productService := service.NewProductService(productRepository)
	supplierService := service.NewSupplierService(supplierRepository)
	transactionService := service.NewTransactionService(transactionRepository, productQualityRepository, txRepository)

	// Init controllers and routes
	prefix := app.Group("/api")
	app.Get("/", WelcomeHandler)

	controller.NewAuthController(userService, prefix)

	app.Use(middleware.NewJWTMiddleware())

	controller.NewUserController(userService, es, prefix)
	controller.NewCustomerController(customerService, prefix)
	controller.NewProductQualityController(productQualityService, prefix)
	controller.NewProductController(productService, prefix)
	controller.NewSupplierController(supplierService, prefix)
	controller.NewTransactionController(transactionService, prefix)

	app.Get("*", NotFoundHandler)
}

func NotFoundHandler(c *fiber.Ctx) error {
	return response.ReturnJSON(c, fiber.StatusNotFound, "the requested resource was not found", nil).Build()
}

func WelcomeHandler(c *fiber.Ctx) error {
	return response.ReturnJSON(c, fiber.StatusOK, "Welcome to the Inventory Management API", nil).Build()
}
