package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	UserId string `gorm:"type:uuid;" json:"userId,omitempty"`
	User   *User  `json:"user,omitempty"`

	SubCategoryId string       `gorm:"type:uuid;" json:"subCategoryId,omitempty"`
	SubCategory   *SubCategory `json:"subCategory,omitempty"`
}

func (c *Product) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

type CreateProductPayload struct {
	Name string `json:"name" validate:"required,min=1,max=15"`
}
