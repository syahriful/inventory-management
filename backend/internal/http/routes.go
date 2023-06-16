package http

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"inventory-management/backend/cmd/config"
	"inventory-management/backend/internal/http/controller"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/internal/service"
	"os"
)

func NewInitializedRoutes(configuration config.Config, logFile *os.File) (*fiber.App, error) {
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
	if configuration.Get("STATE") == "production" {
		app.Use(middleware.NewCSRFMiddleware())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to my Inventory API")
	})

	db, err := config.NewPostgresSQLGorm(configuration)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	api := app.Group("/api")
	{
		users := api.Group("/users")
		{
			users.Get("/", userController.FindAll)
			users.Get("/:id", userController.FindByID)
			users.Post("/", userController.Create)
			users.Patch("/:id", userController.Update)
			users.Delete("/:id", userController.Delete)
		}
	}

	return app, nil
}
