package handlers

import (
	"context"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/category"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/predicate"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetCategories(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 15
	count, err := db.DBClient.Category.Query().Count(context.Background())
	if err != nil {
		return fiber.NewError(500, "unable to count categories")
	}
	total_pages := int(math.Ceil(float64(count) / limit))
	if total_pages == 0 {
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        []interface{}{},
			"limit":       limit,
			"total_pages": total_pages,
			"page":        page,
		})
	}
	if total_pages <= page {
		categories, err := db.DBClient.Category.Query().Limit(limit).Offset((page - 1) * limit).Order(ent.Desc(category.FieldCreatedAt)).WithUser().All(context.Background())
		if err != nil {
			return fiber.NewError(500, "unable to get categories")
		}
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        categories,
			"limit":       limit,
			"total_pages": total_pages,
			"page":        page,
		})
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func GetCategoriesOptions(c *fiber.Ctx) error {
	type C []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	var Options C
	err := db.DBClient.Category.Query().Limit(50).Select(category.FieldID, category.FieldName, category.FieldValue).Scan(context.Background(), &Options)
	if err != nil {
		return fiber.NewError(500, "unable to get options")
	}

	if len(Options) == 0 {
		Options = make(C, 0)
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  Options,
	})

}

func CreateCategory(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*ent.User)
	data := struct {
		Name string `json:"name"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	val := strings.ToLower(strings.Join(strings.Split(data.Name, " "), "-"))
	category, err := db.DBClient.Category.Create().SetName(data.Name).SetValue(val).SetUserID(u.ID).Save(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to create category")
	}
	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  category,
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	u, _ := c.Locals("user").(*ent.User)
	data := struct {
		Name string `json:"name"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	val := strings.Join(strings.Split(data.Name, " "), "-")
	var filter predicate.Category
	if u.Type == "ADMIN" {
		filter = category.ID(uuid.MustParse(id))
	} else {
		filter = category.And(category.HasUserWith(user.ID(u.ID)), category.ID(uuid.MustParse(id)))
	}
	category, err := db.DBClient.Category.Update().Where(filter).SetName(data.Name).SetValue(val).Save(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update category")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  category,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	time.Sleep(time.Second * 5)
	id := c.Params("id")
	u, _ := c.Locals("user").(*ent.User)
	var filter predicate.Category
	if u.Type == "ADMIN" {
		filter = category.ID(uuid.MustParse(id))
	} else {
		filter = category.And(category.HasUserWith(user.ID(u.ID)), category.ID(uuid.MustParse(id)))
	}
	if err := db.DBClient.Category.DeleteOneID(uuid.MustParse(id)).Where(filter).Exec(c.Context()); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "category not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    "null",
	})
}
