package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterProductItemRoutes(r fiber.Router, s *db.Store) {
	h := handlers.NewProductHandler(*s)
	authM := middlewares.NewAuthHandler(s.User)
	router := r.Group("/products")
	router.Use(authM.Authorized)
	router.Get("/:subcategorySlug", h.GetProductItems)
	router.Post("/:subcategorySlug", h.AddProduct)
	router.Delete("/:productId", h.DeleteProduct)
	router.Patch("/:productId", h.UpdateProduct)
}
