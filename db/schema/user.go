package schema

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        string     `gorm:"primary_key;column:id;" json:"id"`
	Name      string     `json:"name"`
	Email     string     `gorm:"unique" json:"email,omitempty"`
	Password  string     `json:"-"`
	Type      string     `json:"type,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Categories []Category `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"categories,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New().String()
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

type RegisterUserPayload struct {
	Name     string `json:"name" validate:"required,min=1,max=15"`
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email,lowercase"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

func CreateNewUserParams(p RegisterUserPayload) *User {
	return &User{
		Name:     p.Name,
		Email:    strings.ToLower(p.Email),
		Password: p.Password,
		Type:     "USER",
	}
}
