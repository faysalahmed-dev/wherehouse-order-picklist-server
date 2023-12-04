package handlers

import (
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productStore     store.ProductStore
	subCategoryStore store.SubCategoryStore
}

func NewProductHandler(s db.Store) *ProductHandler {
	return &ProductHandler{
		productStore:     s.Product,
		subCategoryStore: s.SubCategory,
	}
}

func (h *ProductHandler) GetProductItems(c *fiber.Ctx) error {
	// u, _ := c.Locals("user").(*schema.User)
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	subcategorySlug := c.Params("subcategorySlug")
	if len(subcategorySlug) == 0 {
		return fiber.NewError(400, "sub category required")
	}
	subCat, err := h.subCategoryStore.GetByFields(&schema.SubCategory{Value: subcategorySlug})
	if len(subcategorySlug) == 0 {
		return fiber.NewError(400, "sub category not found")
	}
	const limit = 30

	condition := &schema.Product{SubCategoryId: subCat.ID}
	tP, err := h.productStore.Pagination(limit, condition)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if tP == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error":       false,
			"data":        []interface{}{},
			"limit":       limit,
			"total_pages": tP,
			"page":        page,
		})
	}
	if tP <= page {
		results, err := h.productStore.GetAll(page, limit, condition)
		if err != nil {
			return fiber.NewError(500, "unable to get products")
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

func (h *ProductHandler) AddProduct(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	subcategorySlug := c.Params("subcategorySlug")
	if len(subcategorySlug) == 0 {
		return fiber.NewError(400, "sub category required")
	}

	subCat, err := h.subCategoryStore.GetByFields(&schema.SubCategory{Value: subcategorySlug})
	if len(subcategorySlug) == 0 {
		return fiber.NewError(400, "sub category not found")
	}
	var data schema.CreateProductPayload

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	p, err := h.productStore.InsertOne(&schema.Product{
		Name:          data.Name,
		UserId:        u.ID,
		SubCategoryId: subCat.ID,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to add product")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"data":  p,
	})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	productId, err := uuid.Parse(c.Params("productId"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid product id")
	}
	var data schema.CreateProductPayload
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	var condition *schema.Product

	if u.Type == "ADMIN" {
		condition = &schema.Product{ID: productId.String()}
	} else {
		condition = &schema.Product{ID: productId.String(), UserId: u.ID}
	}
	result, err := h.productStore.UpdateOne(condition, &schema.Product{Name: data.Name})
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "record not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"data":  result,
	})
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*schema.User)
	pId, err := uuid.Parse(c.Params("productId"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if u.Type == "ADMIN" {
		err = h.productStore.DeleteById(pId.String())
	} else {
		err = h.productStore.DeleteByUserAndId(u.ID, pId.String())
	}
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "record not found")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    nil,
	})
}
