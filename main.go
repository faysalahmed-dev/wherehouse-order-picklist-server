package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		AppName:       "Whorehouse Order Picker",
		ErrorHandler:  handlers.GlobalErrorHandler,
	})
	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(helmet.New())
	// app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())
	// app.Get("/stats", monitor.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"error":   false,
			"message": "warehouse api",
		})
	})
	apiV1 := app.Group("/api/v1")
	routes.RegisterUserRoutes(apiV1)
	routes.RegisterCategoriesRoutes(apiV1)
	routes.RegisterSubCategoriesRoutes(apiV1)
	routes.RegisterOrdersRoutes(apiV1)

	go db.ConnectToDB()

	runtimeOs := runtime.GOOS
	port := os.Getenv("PORT")
	fmt.Println("os: ", runtimeOs)
	if runtimeOs == "windows" {
		log.Fatal(app.Listen("127.0.0.1:" + port))
	} else {
		log.Fatal(app.Listen(":" + port))
	}
}
