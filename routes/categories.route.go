package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterCategoriesRoutes(r fiber.Router) {
	router := r.Group("/categories")
	router.Use(middlewares.Authorized)
	router.Get("", handlers.GetCategories)
	router.Get("/options", handlers.GetCategoriesOptions)
	router.Post("/add-category", handlers.CreateCategory)
	router.Patch("/edit/:id", handlers.UpdateCategory)
	router.Delete("/:id", handlers.DeleteCategory)
}
