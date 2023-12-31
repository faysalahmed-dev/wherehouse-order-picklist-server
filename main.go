package main

import (
	"log"
	"os"
	"runtime"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
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

	var (
		dbClient = db.ConnectToDB()
		dbStore  = &db.Store{
			User:        store.NewUserStore(dbClient),
			Category:    store.NewCategoryStore(dbClient),
			SubCategory: store.NewSubCategoryStore(dbClient),
			Product:     store.NewProductStore(dbClient),
			Order:       store.NewOrderStore(dbClient),
		}
	)

	defer func() {
		d, _ := dbClient.DB()
		d.Close()
	}()

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

	apiV1 := app.Group("/api/v1")
	routes.RegisterUserRoutes(apiV1, dbStore)
	routes.RegisterCategoriesRoutes(apiV1, dbStore)
	routes.RegisterSubCategoriesRoutes(apiV1, dbStore)
	routes.RegisterOrdersRoutes(apiV1, dbStore)
	routes.RegisterProductItemRoutes(apiV1, dbStore)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"error":   false,
			"message": "warehouse api",
		})
	})
	app.Get("/all-routes", func(c *fiber.Ctx) error {
		return c.JSON(app.Stack())
	})

	runtimeOs := runtime.GOOS
	port := os.Getenv("PORT")

	if runtimeOs == "windows" {
		log.Fatal(app.Listen("127.0.0.1:" + port))
	} else {
		log.Fatal(app.Listen(":" + port))
	}
}
