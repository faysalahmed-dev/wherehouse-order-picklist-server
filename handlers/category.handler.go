package handlers

import (
	"fmt"
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
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
	tP, err := h.categoryStore.Pagination(limit)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}
	if tP == 0 {
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        []interface{}{},
			"limit":       limit,
			"total_pages": tP,
			"page":        page,
		})
	}
	if tP <= page {
		results, err := h.categoryStore.GetCategories(page, limit)
		if err != nil {
			return fiber.NewError(500, "unable to get categories")
		}
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        results,
			"limit":       limit,
			"total_pages": tP,
			"page":        page,
		})
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func (h *CategoryHandler) GetCategoriesOptions(c *fiber.Ctx) error {
	results, err := h.categoryStore.GetCategoryOptions(1, 50)
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
