package handlers

import (
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SubCategoryHandler struct {
	subCategoryStore store.SubCategoryStore
	categoryStore    store.CategoryStore
}

func NewSubCategoryHandler(s store.SubCategoryStore, cs store.CategoryStore) *SubCategoryHandler {
	return &SubCategoryHandler{
		subCategoryStore: s,
		categoryStore:    cs,
	}
}

func (h *SubCategoryHandler) GetSubCategories(c *fiber.Ctx) error {
	categoryParam := c.Params("category", "")
	if len(categoryParam) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "category required")
	}
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 20
	category, err := h.categoryStore.GetByFields(&schema.Category{Value: categoryParam})
	if err != nil {
		return fiber.NewError(400, "category not found")
	}
	condition := &schema.SubCategory{CategoryId: category.ID}
	tP, err := h.subCategoryStore.Pagination(limit, condition)
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
		results, err := h.subCategoryStore.GetAll(page, limit, condition)
		if err != nil {
			return fiber.NewError(500, "unable to get sub categories")
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

func (h *SubCategoryHandler) GetSubCategoriesOptions(c *fiber.Ctx) error {
	categoryParam := c.Params("category", "")
	if len(categoryParam) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "category required")
	}

	results, err := h.subCategoryStore.GetOptions(1, 50, &schema.SubCategory{Category: &schema.Category{Value: categoryParam}})
	if err != nil {
		return fiber.NewError(500, "unable to get options")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  results,
	})
}

func (h *SubCategoryHandler) CreateSubCategory(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	categoryParam := c.Params("category", "")
	if len(categoryParam) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "category required")
	}
	var data schema.CreateSubCategoryPayload
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}

	category, err := h.categoryStore.GetByFields(&schema.Category{Value: categoryParam})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "category not found")
	}

	sub_category, err := h.subCategoryStore.InsertOne(schema.CreateSubCategoryParams(data, u.ID, category.ID))

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to create sub category")
	}
	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  sub_category,
	})
}

func (h *SubCategoryHandler) UpdateSubCategory(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	id, err := uuid.Parse(c.Params("id", ""))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	var data schema.CreateSubCategoryPayload
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}

	var condition *schema.SubCategory

	if u.Type == "ADMIN" {
		condition = &schema.SubCategory{ID: id.String()}
	} else {
		condition = &schema.SubCategory{ID: id.String(), UserId: u.ID}
	}
	result, err := h.subCategoryStore.UpdateOne(condition, schema.UpdateSubCategoryParams(data))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "record not found")
	}

	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  result,
	})
}

func (h *SubCategoryHandler) DeleteSubCategory(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id", ""))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	u, _ := c.Locals("user").(*schema.User)

	if u.Type == "ADMIN" {
		err = h.subCategoryStore.DeleteById(id.String())
	} else {
		err = h.subCategoryStore.DeleteByUserAndId(u.ID, id.String())
	}
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "record not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    nil,
	})
}
