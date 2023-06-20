package http

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"inventory-management/backend/cmd/config"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/http/routes"
	"os"
)

func NewInitializedRoutes(configuration config.Config, logFile *os.File) (*fiber.App, error) {
	// Init app and middlewares
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return response.ReturnError(ctx, code, err)
		},
	})
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
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to my Inventory API")
	})

	api := app.Group("/api")
	routes.NewUserRoute(db, api)
	routes.NewProductRoute(db, api)
	routes.NewProductQualityRoute(db, api)
	routes.NewSupplierRoute(db, api)

	return app, nil
}
