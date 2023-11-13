package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(r fiber.Router) {
	userRoute := r.Group("/user")
	userRoute.Post("/login", handlers.LoginUser)
	userRoute.Post("/register", handlers.RegisterUser)
	userRoute.Get("/profile", middlewares.Authorized, handlers.Profile)
}
