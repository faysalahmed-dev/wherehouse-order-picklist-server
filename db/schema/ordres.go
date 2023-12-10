package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID       string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Amount   int    `json:"amount"`
	UnitType string `json:"unit"`
	// PENDING || DONE
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`

	UserId string `gorm:"type:uuid;" json:"userId,omitempty"`
	User   *User  `json:"user,omitempty"`

	ProductId string   `gorm:"type:uuid;" json:"productId,omitempty"`
	Product   *Product `json:"product,omitempty"`
}

func (c *Order) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	return nil
}

type GroupOrderByUser struct {
	Total          int    `json:"total"`
	Picked         int    `json:"picked"`
	UserId         string `json:"user_id"`
	LastSubmission string `json:"last_submission"`
	Name           string `json:"name"`
	Email          string `json:"email"`
}

type OrderUserOptions struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type CreateOrderPayload struct {
	Amount   int    `json:"amount" validate:"required,min=1,max=15"`
	UnitType string `json:"unit_type" validate:"required,min=1,max=15"`
}

func CreateOrderParams(c CreateOrderPayload, uId string, pid string) *Order {
	return &Order{
		Amount:    c.Amount,
		UnitType:  c.UnitType,
		Status:    "PENDING",
		UserId:    uId,
		ProductId: pid,
	}
}
