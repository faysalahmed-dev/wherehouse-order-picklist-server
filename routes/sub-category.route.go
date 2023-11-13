package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterSubCategoriesRoutes(r fiber.Router) {
	router := r.Group("/sub-categories")
	router.Use(middlewares.Authorized)
	router.Get("/:category", handlers.GetSubCategories)
	router.Get("/:category/options", handlers.GetSubCategoriesOptions)
	router.Post("/add-subcategory", handlers.CreateSubCategory)
	router.Patch("/edit/:id", handlers.UpdateSubCategory)
	router.Delete("/:id", handlers.DeleteSubCategory)
}
