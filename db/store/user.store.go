package store

import (
	"errors"
	"fmt"
	"strings"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUserById(userId string) (*schema.User, error)
	GetUserByEmail(email string) (*schema.User, error)
	InsertUser(u *schema.User) (*schema.User, error)
}

type DBUserStore struct {
	client *gorm.DB
	model  *gorm.DB
}

func NewUserStore(client *gorm.DB) *DBUserStore {
	return &DBUserStore{
		client: client,
		model:  client.Model(schema.User{}),
	}
}

func (s *DBUserStore) GetUserById(userId string) (*schema.User, error) {
	var user schema.User
	err := s.model.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, errors.New("recode not found")
	}
	return &user, nil
}
func (s *DBUserStore) GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	err := s.model.Where("email = ?", strings.ToLower(email)).First(&user).Error
	if err != nil {
		return nil, errors.New("recode not found")
	}
	return &user, nil
}

func (s *DBUserStore) InsertUser(user *schema.User) (*schema.User, error) {
	fmt.Printf("%+v", *user)
	result := s.client.Create(user)
	fmt.Println("49: ", result.Error)
	if result.Error != nil {
		return nil, errors.New("unable to insert user")
	}
	fmt.Println(user)
	return user, nil
}
