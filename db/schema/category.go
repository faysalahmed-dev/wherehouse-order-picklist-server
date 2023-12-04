package schema

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string     `json:"name"`
	Value     string     `gorm:"unique" json:"value"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	UserId string `gorm:"type:uuid;" json:"userId,omitempty"`
	User   *User  `json:"user,omitempty"`

	SubCategories []SubCategory `gorm:"foreignKey:CategoryId;constraint:OnDelete:CASCADE" json:"subCategories,omitempty"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

type CreateCategoryPayload struct {
	Name string `json:"name" validate:"required,min=1,max=15"`
}

func CreateCategoryParams(c CreateCategoryPayload, userId string) *Category {
	val := strings.ToLower(strings.Join(strings.Split(c.Name, " "), "-"))
	return &Category{
		Name:   c.Name,
		Value:  val,
		UserId: userId,
	}
}

func UpdateCategoryParams(c CreateCategoryPayload) *Category {
	val := strings.ToLower(strings.Join(strings.Split(c.Name, " "), "-"))
	return &Category{
		Name:  c.Name,
		Value: val,
	}
}
