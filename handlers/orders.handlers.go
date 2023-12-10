package handlers

import (
	"fmt"
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderHandler struct {
	orderStore   store.OrderStore
	productStore store.ProductStore
}

func NewOrderHandler(s db.Store) *OrderHandler {
	return &OrderHandler{
		orderStore:   s.Order,
		productStore: s.Product,
	}
}

func (h *OrderHandler) GetOrdersByUser(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "invalid page")
	}
	userId := c.Query("userId", "")
	if len(userId) > 0 {
		_, err := uuid.Parse(userId)
		if err != nil {
			return fiber.NewError(400, "invalid user id")
		}
	}
	const limit = 20
	opt := store.PaginationOpt{
		Page:  page,
		Limit: limit,
	}
	pOpt, err := h.orderStore.Pagination(&schema.Order{UserId: userId}, opt)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	if pOpt.TotalPages == 0 {
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, []interface{}{})
	}
	if pOpt.PageNum <= page {
		results, err := h.orderStore.GroupOrderByUser(userId, opt)
		if err != nil {
			return fiber.NewError(500, "unable to get get orders")
		}
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, results)
	} else {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}

func (h *OrderHandler) GetUserOrders(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("userId", ""))
	if err != nil {
		return fiber.NewError(400, "invalid user id")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 20
	opt := store.PaginationOpt{
		Page:  page,
		Limit: limit,
	}
	pOpt, err := h.orderStore.Pagination(&schema.Order{UserId: userId.String()}, opt)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	if pOpt.TotalPages == 0 {
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, []interface{}{})
	}
	if pOpt.TotalPages <= page {
		results, err := h.orderStore.OrdersByUserId(userId.String(), opt)
		if err != nil {
			return fiber.NewError(500, "unable to get orders")
		}
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, results)

	} else {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}

func (h *OrderHandler) MyOrders(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 20
	opt := store.PaginationOpt{
		Page:  page,
		Limit: limit,
	}
	pOpt, err := h.orderStore.Pagination(&schema.Order{UserId: u.ID}, opt)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	if pOpt.PageNum == 0 {
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, []interface{}{})
	}
	if pOpt.PageNum <= page {
		results, err := h.orderStore.OrdersByUserId(u.ID, opt)
		if err != nil {
			return fiber.NewError(500, "unable to get categories")
		}
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pOpt, Limit: limit}, results)
	} else {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}

func (h *OrderHandler) UserOptions(c *fiber.Ctx) error {
	results, err := h.orderStore.HasOrderUsers(store.PaginationOpt{
		Page:  1,
		Limit: 50,
	})
	if err != nil {
		return fiber.NewError(500, "unable to get options")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  results,
	})
}

func (h *OrderHandler) AddOrder(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	pId, err := uuid.Parse(c.Params("productId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid product id")
	}
	p, err := h.productStore.GetByFields(&schema.Product{ID: pId.String()})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "product not found")
	}
	if h.orderStore.HasOrder(p.ID, u.ID) {
		return fiber.NewError(fiber.StatusBadRequest, "order already exists")
	}
	var data schema.CreateOrderPayload
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	newO, err := h.orderStore.InsertOne(schema.CreateOrderParams(data, u.ID, p.ID))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to create order")
	}
	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  newO,
	})
}

func (h *OrderHandler) UpdateOrder(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	oId, err := uuid.Parse(c.Params("orderId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid params id")
	}

	var data *schema.CreateOrderPayload

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	var result *schema.Order
	if u.Type == "ADMIN" {
		result, err = h.orderStore.UpdateByFields(&schema.Order{ID: oId.String()}, &schema.Order{Amount: data.Amount, UnitType: data.UnitType})
	} else {
		result, err = h.orderStore.UpdateByFields(&schema.Order{ID: oId.String(), UserId: u.ID}, &schema.Order{Amount: data.Amount, UnitType: data.UnitType})
	}
	fmt.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "order not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  result,
	})
}

func (h *OrderHandler) UpdateStatus(c *fiber.Ctx) error {
	oId, err := uuid.Parse(c.Params("orderId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid params id")
	}

	data := struct {
		Status string `json:"status"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}

	result, err := h.orderStore.UpdateByFields(&schema.Order{ID: oId.String()}, &schema.Order{
		Status: data.Status,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "order not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  result,
	})
}

func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	oid, err := uuid.Parse(c.Params("orderId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid order id")
	}

	if u.Type == "ADMIN" {
		err = h.orderStore.DeleteByFields(&schema.Order{ID: oid.String()})
	} else {
		err = h.orderStore.DeleteByFields(&schema.Order{ID: oid.String(), UserId: u.ID})
	}

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "order not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  nil,
	})
}
