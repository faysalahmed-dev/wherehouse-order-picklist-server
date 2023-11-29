package handlers

import (
	"fmt"
	"math"
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	_, err := db.DBClient.User.Query().Where(user.Email(userData.Email)).First(c.Context())
	if err == nil {
		return fiber.NewError(fiber.StatusConflict, "user already exits")
	}

	hashPass, err := helpers.HashPassword(userData.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to hash password")
	}
	newUser, err := db.DBClient.User.Create().SetEmail(userData.Email).SetPassword(hashPass).SetName(userData.Name).Save(c.Context())
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

	user, err := db.DBClient.User.Query().Where(user.Email(userData.Email)).First(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "user not found")
	}
	if user.Blocked {
		return fiber.NewError(fiber.StatusForbidden, "account blocked")
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

func GetAllUser(c *fiber.Ctx) error {
	ctx := c.Context()
	const limit = 15
	status_type := c.Query("status_type") // "all" || "blocked" || "unblocked"
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "invalid page number")
	}
	filters := user.TypeEQ("USER")
	if status_type == "blocked" {
		filters = user.And(filters, user.Blocked(true))
	} else if status_type == "unblocked" {
		filters = user.And(filters, user.Blocked(false))
	}

	count, err := db.DBClient.User.Query().Where(filters).Count(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to count")
	}
	total_pages := int(math.Ceil(float64(count) / limit))
	if total_pages == 0 {
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        []interface{}{},
			"limit":       limit,
			"total_pages": total_pages,
			"page":        page,
			"total_items": count,
		})
	}
	if page <= total_pages {
		users, err := db.DBClient.
			User.
			Query().
			Where(filters).
			Limit(limit).
			Order(ent.Desc(user.FieldCreatedAt)).
			Offset((page - 1) * limit).
			All(ctx)
		if err != nil {
			return fiber.NewError(500, "unable to get users")
		}
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        users,
			"limit":       limit,
			"total_pages": total_pages,
			"page":        page,
			"total_items": count,
		})
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func SearchUserByName(c *fiber.Ctx) error {
	ctx := c.Context()
	status_type := c.Query("status_type") // "all" || "blocked" || "unblocked"
	name := c.Query("name")
	if len(name) == 0 {
		return fiber.NewError(400, "name query can not be empty")
	}
	filters := user.TypeEQ("USER")
	if status_type == "blocked" {
		filters = user.And(filters, user.Blocked(true))
	} else if status_type == "unblocked" {
		filters = user.And(filters, user.Blocked(false))
	}
	users, err := db.DBClient.
		User.
		Query().
		Where(filters).
		Where(user.NameContains(name)).
		Limit(5).
		Order(ent.Desc(user.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return fiber.NewError(500, "unable to get users")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := db.DBClient.User.DeleteOneID(uuid.MustParse(id)).Exec(c.Context()); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "user not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    "null",
	})
}

func UpdateUserStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Println(string(c.Body()))
	data := struct {
		Status bool `json:"status"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid payload")
	}
	user, err := db.DBClient.User.Query().Where(user.ID(uuid.MustParse(id))).First(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	if user.Type.String() == "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "admin account can't be blocked")
	}
	if _, err := db.DBClient.User.UpdateOneID(user.ID).SetBlocked(data.Status).Save(c.Context()); err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update status")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully updated the status",
	})
}
