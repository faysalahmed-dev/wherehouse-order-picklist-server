package handlers

import (
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	userStore store.UserStore
}

func NewUserHandler(userStore store.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var userParams schema.RegisterUserPayload
	if err := c.BodyParser(&userParams); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	if len(userParams.Password) == 0 || len(userParams.Email) == 0 || len(userParams.Name) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}
	_, err := h.userStore.GetUserByEmail(userParams.Email)
	if err == nil {
		return fiber.NewError(fiber.StatusConflict, "user already exits")
	}

	newUser, err := h.userStore.InsertUser(schema.CreateNewUserParams(userParams))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable create user")
	}
	token, err := helpers.GenToken(newUser.ID)
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

func (h *UserHandler) LoginUser(c *fiber.Ctx) error {
	var userParams schema.LoginUserPayload
	if err := c.BodyParser(&userParams); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	if len(userParams.Password) == 0 || len(userParams.Email) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	}

	user, err := h.userStore.GetUserByEmail(userParams.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusForbidden, "user not found")
	}
	if user.Blocked == 1 {
		return fiber.NewError(fiber.StatusForbidden, "account blocked")
	}

	if !helpers.CheckPasswordHash(userParams.Password, user.Password) {
		return fiber.NewError(fiber.StatusForbidden, "password not match")
	}
	token, err := helpers.GenToken(user.ID)
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

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*schema.User)
	if ok {
		return c.Status(200).JSON(fiber.Map{
			"error": false,
			"data":  u,
		})
	}

	return fiber.NewError(fiber.StatusInternalServerError, "unable to find user")
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	uId, err := uuid.Parse(c.Params("userId", ""))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}
	u, err := h.userStore.GetUserById(uId.String())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  u,
	})
}

func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	status_type := c.Query("status_type") // "all" || "blocked"
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "invalid page number")
	}
	var filters *schema.User
	if status_type == "blocked" {
		filters = &schema.User{Blocked: 1, Type: "USER"}
	} else {
		filters = &schema.User{Blocked: 0, Type: "USER"}
	}

	const limit = 20
	opt := store.PaginationOpt{Limit: limit, Page: page}
	pOpt, err := h.userStore.Pagination(filters, opt)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	users, err := h.userStore.GetAll(filters, opt)
	if pOpt.TotalPages == 0 {
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, []interface{}{})
	}
	if pOpt.PageNum <= page {
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, users)
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user id")
	}
	user, err := h.userStore.GetUserById(id.String())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	if user.Type == "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "admin account can't be deleted")
	}
	if err := h.userStore.DeleteById(id.String()); err != nil {
		return fiber.NewError(fiber.StatusForbidden, "unable to delete user")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    nil,
	})
}

func (h *UserHandler) UpdateUserStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	data := struct {
		Blocked int `json:"blocked"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid payload")
	}
	user, err := h.userStore.GetUserById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "user not found")
	}
	if user.Type == "ADMIN" {
		return fiber.NewError(fiber.StatusForbidden, "admin account can't be blocked")
	}
	_, err = h.userStore.UpdateById(user.ID, map[string]interface{}{"blocked": data.Blocked})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update status")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully updated the status",
	})
}
