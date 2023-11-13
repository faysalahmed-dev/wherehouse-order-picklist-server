package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterOrdersRoutes(r fiber.Router) {
	router := r.Group("/orders")
	router.Use(middlewares.Authorized)
	router.Get("", handlers.GetOrders)
	router.Get("/options/:sub_category", handlers.GetOrdersOptions)
	router.Get("/pick-list", middlewares.AdminOnly, handlers.GetPickList)
	router.Post("/add-order", handlers.AddOrders)
	router.Patch("/edit/:id", handlers.UpdateOrder)
	router.Patch("/update-status/:id", middlewares.AdminOnly, handlers.AddToPickList)

	router.Delete("/:id", handlers.DeleteOrder)
}
