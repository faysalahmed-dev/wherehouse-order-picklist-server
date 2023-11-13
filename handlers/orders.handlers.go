package handlers

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/category"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/order"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/predicate"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/subcategory"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/user"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetOrders(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*ent.User)
	page, err := strconv.Atoi(c.Query("page", "1"))
	qC := c.Query("category")
	qSubC := c.Query("sub_category")

	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	const limit = 15
	var filters predicate.Order
	if u.Type == "ADMIN" {
		filters = order.NameNEQ("")
	} else {
		filters = order.HasUserWith(user.ID(u.ID))
	}
	if len(qC) > 0 {
		filters = order.And(filters, order.HasSubCategoriesWith(subcategory.HasCategoryWith(category.ValueContains(qC))))
	}
	if len(qSubC) > 0 {
		filters = order.And(filters, order.HasSubCategoriesWith(subcategory.ValueContains(qSubC)))
	}
	count, err := db.DBClient.Order.Query().Where(filters).Count(context.Background())
	if err != nil {
		return fiber.NewError(500, "unable to count orders")
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
		orders, err := db.DBClient.Order.Query().
			Limit(limit).
			Where(filters).
			WithSubCategories(func(scq *ent.SubCategoryQuery) {
				// var v struct {
				// 	ID           string `json:"id"`
				// 	Name         string `json:"name"`
				// 	Descriptions string `json:"descriptions"`
				// 	Value        string `json:"value"`
				// }
				scq.Select(subcategory.FieldID, subcategory.FieldName, subcategory.FieldDescriptions, subcategory.FieldValue)
				scq.WithCategory().Select(category.FieldID, category.FieldName, category.FieldValue)
			}).
			WithUser(func(uq *ent.UserQuery) {
				uq.Select(user.FieldName, user.FieldType, user.FieldID)
			}).
			Order(ent.Desc(order.FieldCreatedAt)).
			Offset((page - 1) * limit).
			All(context.Background())

		if err != nil {
			return fiber.NewError(500, "unable to get orders")
		}
		return c.Status(200).JSON(fiber.Map{
			"error":       false,
			"data":        orders,
			"limit":       limit,
			"total_pages": total_pages,
			"total_items": count,
			"page":        page,
		})
	} else {
		return fiber.NewError(404, "page limit exit")
	}
}

func GetPickList(c *fiber.Ctx) error {
	status := c.Query("status")
	page, err := strconv.Atoi(c.Query("page", "1"))
	qC := c.Query("category")
	qSubC := c.Query("sub_category")
	if err != nil {
		return fiber.NewError(400, "page num is invalid")
	}
	fmt.Println(status == order.StatusPICKED.String() || status == order.StatusUNPICKED.String())
	if status == order.StatusPICKED.String() || status == order.StatusUNPICKED.String() {
		const limit = 15
		var filters predicate.Order

		filters = order.StatusEQ(order.Status(status))
		if len(qC) > 0 {
			filters = order.And(filters, order.HasSubCategoriesWith(subcategory.HasCategoryWith(category.ValueContains(qC))))
		}
		if len(qSubC) > 0 {
			filters = order.And(filters, order.HasSubCategoriesWith(subcategory.ValueContains(qSubC)))
		}
		count, err := db.DBClient.Order.Query().Where(filters).Count(context.Background())
		if err != nil {
			return fiber.NewError(500, "unable to count orders")
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
			orders, err := db.DBClient.Order.Query().
				Limit(limit).
				Where(filters).
				WithSubCategories(func(scq *ent.SubCategoryQuery) {
					// var v struct {
					// 	ID           string `json:"id"`
					// 	Name         string `json:"name"`
					// 	Descriptions string `json:"descriptions"`
					// 	Value        string `json:"value"`
					// }
					scq.Select(subcategory.FieldID, subcategory.FieldName, subcategory.FieldDescriptions, subcategory.FieldValue)
					scq.WithCategory().Select(category.FieldID, category.FieldName, category.FieldValue)
				}).
				WithUser(func(uq *ent.UserQuery) {
					uq.Select(user.FieldName, user.FieldType, user.FieldID)
				}).
				Order(ent.Desc(order.FieldCreatedAt)).
				Offset((page - 1) * limit).
				All(context.Background())

			if err != nil {
				return fiber.NewError(500, "unable to get orders")
			}
			return c.Status(200).JSON(fiber.Map{
				"error":       false,
				"data":        orders,
				"limit":       limit,
				"total_pages": total_pages,
				"total_items": count,
				"page":        page,
			})
		} else {
			return fiber.NewError(404, "page limit exit")
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("status must be %v or %v", order.StatusPICKED.String(), order.StatusUNPICKED.String()))
}

func GetOrdersOptions(c *fiber.Ctx) error {
	sub_category := c.Params("sub_category")
	type O []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	var Options O
	err := db.DBClient.Order.Query().
		Limit(50).
		Where(order.HasSubCategoriesWith(subcategory.Value(sub_category))).
		Select(order.FieldID, order.FieldName).
		Scan(context.Background(), &Options)

	if err != nil {
		return fiber.NewError(500, "unable to get options")
	}
	if len(Options) == 0 {
		Options = make(O, 0)
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  Options,
	})
}

func AddOrders(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*ent.User)

	data := struct {
		Name        string `json:"name"`
		Amount      string `json:"amount"`
		UnitType    string `json:"unit_type"`
		SubCategory string `json:"sub_category"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	subCategory, err := db.DBClient.SubCategory.Query().Where(subcategory.Value(data.SubCategory)).First(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "sub category not found")
	}
	orderItems, err := db.DBClient.Order.Create().SetName(data.Name).SetAmount(data.Amount).SetUnitType(data.UnitType).SetSubCategoriesID(subCategory.ID).SetUserID(u.ID).Save(context.Background())
	fmt.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to create order")
	}
	return c.Status(201).JSON(fiber.Map{
		"error": false,
		"data":  orderItems,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*ent.User)
	orderid := c.Params("id")
	data := struct {
		Name     string `json:"name"`
		Amount   string `json:"amount"`
		UnitType string `json:"unit_type"`
	}{}
	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}
	var filter predicate.Order

	if u.Type == "ADMIN" {
		filter = order.ID(uuid.MustParse(orderid))
	} else {
		filter = order.And(order.ID(uuid.MustParse(orderid)), order.HasUserWith(user.ID(u.ID)))
	}
	o, err := db.DBClient.Order.Query().Where(filter).First(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "order not found")
	}
	orderItem, err := db.DBClient.Order.Update().Where(order.ID(o.ID)).SetName(data.Name).SetAmount(data.Amount).SetUnitType(data.UnitType).Save(context.Background())
	fmt.Println(err)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update order")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  orderItem,
	})
}

func AddToPickList(c *fiber.Ctx) error {
	orderid := c.Params("id")
	data := struct {
		Status int `json:"status"`
	}{}

	if err := c.BodyParser(&data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "unable to parse body data")
	}

	var status order.Status
	if data.Status == 1 {
		status = order.StatusPICKED
	} else {
		status = order.StatusUNPICKED
	}

	orderItem, err := db.DBClient.Order.Update().Where(order.ID(uuid.MustParse(orderid))).SetStatus(status).Save(context.Background())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to update status")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  orderItem,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	u, _ := c.Locals("user").(*ent.User)
	orderid := c.Params("id")

	var filter predicate.Order

	if u.Type == "ADMIN" {
		filter = order.ID(uuid.MustParse(orderid))
	} else {
		filter = order.And(order.ID(uuid.MustParse(orderid)), order.HasUserWith(user.ID(u.ID)))
	}

	orderItem, err := db.DBClient.Order.Delete().Where(filter).Exec(context.Background())
	if err != nil || orderItem == 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "unable to delete order")
	}
	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"data":  orderItem,
	})
}
