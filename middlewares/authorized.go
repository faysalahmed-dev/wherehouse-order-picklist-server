package middlewares

import (
	"strings"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Authorized(c *fiber.Ctx) error {
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
	user, err := db.DBClient.User.Query().Where(user.ID(user_id)).First(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "user not found")
	}
	if user.Blocked {
		return fiber.NewError(fiber.StatusForbidden, "account blocked")
	}
	c.Locals("user", user)

	return c.Next()
}

func AdminOnly(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*ent.User)

	if u.Type != "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "permission denied")
	}
	return c.Next()
}
