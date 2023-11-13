package handlers

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/category"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/predicate"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/subcategory"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetSubCategories(c *fiber.Ctx) error {
	categoryParam := c.Params("category", "")
	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 15
	filters := subcategory.HasCategoryWith(category.Value(categoryParam))
	count, err := db.DBClient.SubCategory.Query().Where(filters).Count(context.Background())
	if err != nil {
		return fiber.NewError(500, "unable to count sub categories")
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
	if page <= total_pages {
		subcategories, err := db.DBClient.SubCategory.Query().Where(filters).WithUser().Limit(limit).Order(ent.Desc(subcategory.FieldCreatedAt)).Offset((page - 1) * limit).All(context.Background())
		if err != nil {
			return fiber.NewError(500, "unable to get sub categories")
		}
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        subcategories,
			"limit":       limit,
			"total_pages": total_pages,
			"page":        page,
		})
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func GetSubCategoriesOptions(c *fiber.Ctx) error {
	type C []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	var Options C

	categoryParam := c.Params("category", "")

	filters := subcategory.HasCategoryWith(category.Value(categoryParam))

	err := db.DBClient.SubCategory.Query().Where(filters).Limit(50).Select(subcategory.FieldID, subcategory.FieldName, subcategory.FieldValue).Scan(context.Background(), &Options)
	if err != nil {
		return fiber.NewError(500, "unable to get options")
	}
	if len(Options) == 0 {
		Options = make(C, 0)
	}
	// fmt.Println("sub call: ", Options)
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  Options,
	})
}

func CreateSubCategory(c *fiber.Ctx) error {
	data := struct {
		Name         string `json:"name"`
		Descriptions string `json:"descriptions"`
		CategorySlug string `json:"category_slug"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}

	category, err := db.DBClient.Category.Query().Where(category.Value(data.CategorySlug)).First(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "category not found")
	}
	val := strings.ToLower(strings.Join(strings.Split(data.Name, " "), "-"))
	u, _ := c.Locals("user").(*ent.User)

	sub_category, err := db.DBClient.SubCategory.Create().SetName(data.Name).SetValue(val).SetDescriptions(data.Descriptions).SetCategoryID(category.ID).SetUserID(u.ID).Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "unable to create sub category")
	}
	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  sub_category,
	})
}

func UpdateSubCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	u, _ := c.Locals("user").(*ent.User)
	data := struct {
		Name         string `json:"name",omitempty`
		Descriptions string `json:"descriptions",omitempty`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	var filter predicate.SubCategory
	if u.Type == "ADMIN" {
		filter = subcategory.ID(uuid.MustParse(id))
	} else {
		filter = subcategory.And(subcategory.HasUserWith(user.ID(u.ID)), subcategory.ID(uuid.MustParse(id)))
	}
	hasItem, err := db.DBClient.SubCategory.Query().Where(filter).First(context.Background())
	fmt.Println(hasItem)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "item not found")
	}
	var name, descriptions, value string
	if len(data.Name) == 0 {
		name = hasItem.Name
		value = hasItem.Value
	} else {
		name = data.Name
		value = strings.ToLower(strings.Join(strings.Split(data.Name, " "), "-"))
	}
	if len(data.Descriptions) == 0 {
		descriptions = hasItem.Descriptions
	} else {
		descriptions = data.Descriptions
	}

	sub_category, err := db.DBClient.SubCategory.UpdateOneID(hasItem.ID).SetName(name).SetValue(value).SetDescriptions(descriptions).Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update sub category")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  sub_category,
	})
}

func DeleteSubCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	u, _ := c.Locals("user").(*ent.User)
	var filter predicate.SubCategory
	if u.Type == "ADMIN" {
		filter = subcategory.ID(uuid.MustParse(id))
	} else {
		filter = subcategory.And(subcategory.HasUserWith(user.ID(u.ID)), subcategory.ID(uuid.MustParse(id)))
	}
	if err := db.DBClient.SubCategory.DeleteOneID(uuid.MustParse(id)).Where(filter).Exec(c.Context()); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "subcategory not found")
	}
	return c.Status(200).JSON(fiber.Map{
		"error":   false,
		"message": "successfully deleted",
		"data":    "null",
	})
}
