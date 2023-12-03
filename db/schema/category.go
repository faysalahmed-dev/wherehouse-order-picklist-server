package schema

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `json:"name"`
	Value     string    `gorm:"unique" json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryPayload struct {
	Name string `json:"name" validate:"required,min=1,max=15"`
}

func CreateCategoryParams(c CreateCategoryPayload) (*Category, error) {
	val := strings.ToLower(strings.Join(strings.Split(c.Name, " "), "-"))
	return &Category{
		Name:  c.Name,
		Value: val,
	}, nil
}
