package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterCategoriesRoutes(r fiber.Router, s *store.CategoryStore, u *store.UserStore) {
	h := handlers.NewCategoryHandler(*s)
	authM := middlewares.NewAuthHandler(*u)
	router := r.Group("/categories")
	router.Use(authM.Authorized)
	router.Get("", h.GetCategories)
	router.Post("/add-category", h.CreateCategory)
	router.Get("/options", h.GetCategoriesOptions)
	router.Patch("/edit/:id", h.UpdateCategory)
	router.Delete("/:id", h.DeleteCategory)
}
