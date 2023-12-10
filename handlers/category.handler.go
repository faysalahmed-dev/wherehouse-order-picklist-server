package handlers

import (
	"fmt"
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/helpers"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	categoryStore store.CategoryStore
}

func NewCategoryHandler(c store.CategoryStore) *CategoryHandler {
	return &CategoryHandler{
		categoryStore: c,
	}
}

func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 20
	pOpt := store.PaginationOpt{Limit: limit, Page: page}
	pageInfo, err := h.categoryStore.Pagination(&schema.Category{}, pOpt)
	fmt.Println(pageInfo)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	if pageInfo.TotalPages == 0 {
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pageInfo, Limit: limit}, []interface{}{})
	}
	if pageInfo.TotalPages <= page {
		results, err := h.categoryStore.GetCategories(pOpt)
		if err != nil {
			return fiber.NewError(500, "unable to get categories")
		}
		return helpers.SendPaginationRes(c, &helpers.P{PaginationValue: *pageInfo, Limit: limit}, results)
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func (h *CategoryHandler) GetCategoriesOptions(c *fiber.Ctx) error {
	pOpt := store.PaginationOpt{Limit: 50, Page: 1}
	results, err := h.categoryStore.GetCategoryOptions(pOpt)
	if err != nil {
		return fiber.NewError(500, "unable to get options")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  results,
	})
}

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	var data schema.CreateCategoryPayload
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	d := schema.CreateCategoryParams(data, u.ID)
	if _, err := h.categoryStore.GetByFields(&schema.Category{Value: d.Value}); err == nil {
		return fiber.NewError(fiber.StatusConflict, "record already exist")
	}
	category, err := h.categoryStore.InsertCategory(d)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to create category")
	}
	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  category,
	})
}

func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	u, _ := c.Locals("user").(*schema.User)
	var data schema.CreateCategoryPayload
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	newData := schema.UpdateCategoryParams(data)
	_, err := h.categoryStore.GetByFields(&schema.Category{Value: newData.Value})
	if err == nil {
		return fiber.NewError(fiber.StatusConflict, "record already exists")
	}
	var result *schema.Category

	if u.Type == "ADMIN" {
		result, err = h.categoryStore.GetByFields(&schema.Category{ID: id})
	} else {
		result, err = h.categoryStore.GetByFields(&schema.Category{ID: id, UserId: u.ID})
	}
	fmt.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "record not found")
	}
	_, err = h.categoryStore.UpdateById(result.ID, newData)
	fmt.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update record")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  1,
	})
}

func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	u, _ := c.Locals("user").(*schema.User)
	var err error
	if u.Type == "ADMIN" {
		err = h.categoryStore.DeleteById(id)
	} else {
		err = h.categoryStore.DeleteByUserAndId(u.ID, id)
	}
	fmt.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "category not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    nil,
	})
}
