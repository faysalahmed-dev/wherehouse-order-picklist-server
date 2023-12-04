package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterSubCategoriesRoutes(r fiber.Router, s *db.Store) {
	h := handlers.NewSubCategoryHandler(s.SubCategory, s.Category)
	authM := middlewares.NewAuthHandler(s.User)

	router := r.Group("/sub-categories")
	router.Use(authM.Authorized)
	router.Get("/:category", h.GetSubCategories)
	router.Get("/:category/options", h.GetSubCategoriesOptions)
	router.Post("/add-subcategory/:category", h.CreateSubCategory)
	router.Patch("/edit/:id", h.UpdateSubCategory)
	router.Delete("/:id", h.DeleteSubCategory)
}
