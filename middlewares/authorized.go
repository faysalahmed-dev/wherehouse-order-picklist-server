package middlewares

import (
	"strings"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthMiddleware struct {
	userStore store.UserStore
}

func NewAuthHandler(userStore store.UserStore) *AuthMiddleware {
	return &AuthMiddleware{
		userStore: userStore,
	}
}

func (h *AuthMiddleware) Authorized(c *fiber.Ctx) error {
	token := c.Get("authorization", "")
	if len(token) > 0 && strings.Contains(token, "Bearer ") {
		t := strings.Split(token, " ")
		if len(t) == 2 {
			token = t[1]
		} else {
			token = ""
		}
	}
	if len(token) == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "token not found")
	}
	decoded, err := helpers.VerifyToken(token)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid token")
	}

	user_id, err := uuid.Parse(decoded.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user not found")
	}
	user, err := h.userStore.GetUserById(user_id.String())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user not found")
	}
	// if user.Blocked {
	// 	return fiber.NewError(fiber.StatusForbidden, "account blocked")
	// }
	c.Locals("user", user)

	return c.Next()
}

func (h *AuthMiddleware) AdminOnly(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)

	if u.Type != "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "permission denied")
	}
	return c.Next()
}
