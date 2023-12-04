package routes

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/handlers"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(r fiber.Router, s *db.Store) {
	h := handlers.NewUserHandler(s.User)
	userRoute := r.Group("/user")
	userRoute.Post("/login", h.LoginUser)
	userRoute.Post("/register", h.RegisterUser)

	authM := middlewares.NewAuthHandler(s.User)
	userRoute.Use(authM.Authorized)
	userRoute.Get("/profile", h.Profile)
	userRoute.Use(authM.AdminOnly)
	// userRoute.Get("/users", handlers.GetAllUser)
	// userRoute.Get("/search-users", handlers.SearchUserByName)
	// userRoute.Group("/:id").Patch("", handlers.UpdateUserStatus).Delete("", handlers.DeleteUser)
}
