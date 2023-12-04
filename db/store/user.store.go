package store

import (
	"errors"
	"strings"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserStore interface {
	GetUserById(userId string) (*schema.User, error)
	GetUserByEmail(email string) (*schema.User, error)
	InsertUser(u *schema.User) (*schema.User, error)
}

type DBUserStore struct {
	client *gorm.DB
}

func NewUserStore(client *gorm.DB) *DBUserStore {
	return &DBUserStore{
		client: client,
	}
}

func (s *DBUserStore) GetUserById(userId string) (*schema.User, error) {
	var user schema.User
	err := s.client.Model(&schema.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, errors.New("record not found")
	}
	return &user, nil
}
func (s *DBUserStore) GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	err := s.client.Model(&schema.User{}).Where("email = ?", strings.ToLower(email)).First(&user).Error
	if err != nil {
		return nil, errors.New("record not found")
	}
	return &user, nil
}

func (s *DBUserStore) InsertUser(user *schema.User) (*schema.User, error) {
	result := s.client.Omit(clause.Associations).Create(user)
	if result.Error != nil {
		return nil, errors.New("unable to insert user")
	}
	return user, nil
}
