package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterOrdersRoutes(r fiber.Router, s *db.Store) {
	h := handlers.NewOrderHandler(*s)
	authM := middlewares.NewAuthHandler(s.User)
	router := r.Group("/orders")
	router.Use(authM.Authorized)

	router.Post("/:productId/add-order", h.AddOrder)
	router.Get("/my-orders", h.MyOrders)
	router.Delete("/:orderId", h.DeleteOrder)
	router.Patch("/edit/:orderId", h.UpdateOrder)

	router.Use(authM.AdminOnly)
	router.Get("", h.GetOrdersByUser)
	router.Get("/user-options", h.UserOptions)
	router.Get("/view-orders/:userId", h.GetUserOrders)
	router.Patch("/update-status/:orderId", authM.AdminOnly, h.UpdateStatus)
	// router.Get("/options/:sub_category", handlers.GetOrdersOptions)
	// router.Get("/pick-list", middlewares.AdminOnly, handlers.GetPickList)

}
