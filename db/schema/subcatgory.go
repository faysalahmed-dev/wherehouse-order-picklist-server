package schema

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubCategory struct {
	ID          string     `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string     `json:"name"`
	Value       string     `json:"value"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`

	UserId string `gorm:"type:uuid;" json:"userId,omitempty"`
	User   *User  `json:"user,omitempty"`

	CategoryId string    `gorm:"type:uuid;" json:"categoryId,omitempty"`
	Category   *Category `json:"category,omitempty"`
}

func (c *SubCategory) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

type CreateSubCategoryPayload struct {
	Name        string `json:"name" validate:"required,min=1,max=15"`
	Description string `json:"description" validate:"required,min=1,max=15"`
}

func CreateSubCategoryParams(c CreateSubCategoryPayload, uId string, cid string) *SubCategory {
	val := strings.ToLower(strings.Join(strings.Split(c.Name, " "), "-"))
	return &SubCategory{
		Name:        c.Name,
		Value:       val,
		Description: c.Description,
		UserId:      uId,
		CategoryId:  cid,
	}
}

func UpdateSubCategoryParams(c CreateSubCategoryPayload) *SubCategory {
	val := strings.ToLower(strings.Join(strings.Split(c.Name, " "), "-"))
	return &SubCategory{
		Name:        c.Name,
		Value:       val,
		Description: c.Description,
	}
}
