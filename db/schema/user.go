package schema

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"-"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func CreateNewUserParams(p RegisterUserPayload) (*User, error) {
	// err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	bytes, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:     p.Name,
		Email:    p.Email,
		Password: string(bytes),
		Type:     "USER",
	}, nil
}

func ComparePassword(userPassword string, savedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(savedPassword), []byte(userPassword))
}
