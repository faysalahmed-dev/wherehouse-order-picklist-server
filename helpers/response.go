package helpers

import (
	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/store"
	"github.com/gofiber/fiber/v2"
)

type P struct {
	store.PaginationValue
	Limit int
}

func SendPaginationRes(c *fiber.Ctx, p *P, data any) error {
	return c.Status(200).JSON(fiber.Map{
		"error":       false,
		"data":        data,
		"limit":       p.Limit,
		"total_pages": p.TotalPages,
		"page":        p.PageNum,
		"total_items": p.TotalItems,
	})
}
