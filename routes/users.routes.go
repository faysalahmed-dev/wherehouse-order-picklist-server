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

	userRoute.Use(middlewares.Authorized)
	userRoute.Get("/profile", handlers.Profile)
	userRoute.Use(middlewares.AdminOnly)
	userRoute.Get("/users", handlers.GetAllUser)
	userRoute.Get("/search-users", handlers.SearchUserByName)
	userRoute.Group("/:id").Patch("", handlers.UpdateUserStatus).Delete("", handlers.DeleteUser)
}
