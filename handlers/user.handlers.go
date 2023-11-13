package handlers

import (
	"context"
	"fmt"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	userData := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&userData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	if len(userData.Password) == 0 || len(userData.Email) == 0 || len(userData.Name) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	_, err := db.DBClient.User.Query().Where(user.Email(userData.Email)).First(context.Background())
	if err == nil {
		return fiber.NewError(fiber.StatusConflict, "user already exits")
	}

	hashPass, err := helpers.HashPassword(userData.Password)
	fmt.Println("hash pass: ", hashPass)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to hash password")
	}
	newUser, err := db.DBClient.User.Create().SetEmail(userData.Email).SetPassword(hashPass).SetName(userData.Name).Save(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable create user")
	}
	token, err := helpers.GenToken(newUser.ID.String())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable gen token")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully user created",
		"data": fiber.Map{
			"token": token,
			"user":  newUser,
		},
	})
}

func LoginUser(c *fiber.Ctx) error {
	userData := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&userData); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	if len(userData.Password) == 0 || len(userData.Email) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	user, err := db.DBClient.User.Query().Where(user.Email(userData.Email)).First(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "user not found")
	}

	if !helpers.CheckPasswordHash(userData.Password, user.Password) {
		return fiber.NewError(fiber.StatusForbidden, "password not match")
	}
	token, err := helpers.GenToken(user.ID.String())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable gen token")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully user created",
		"data": fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}

func Profile(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)
	if ok {
		return c.Status(200).JSON(fiber.Map{
			"error": false,
			"data":  u,
		})
	}

	return fiber.NewError(fiber.StatusInternalServerError, "unable to find user")
}

// func TestFunc(c *fiber.Ctx) error {
// 	u, ok := c.Locals("user").(*ent.User)
// 	fmt.Println(ok)
// 	fmt.Println(u)
// 	return c.Status(200).JSON(fiber.Map{"user_id": u.ID})
// }
