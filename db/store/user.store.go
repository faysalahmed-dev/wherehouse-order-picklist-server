package store

import (
	"errors"
	"math"
	"strings"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserStore interface {
	Pagination(condition *schema.User, opt PaginationOpt) (*PaginationValue, error)
	GetUserById(userId string) (*schema.User, error)
	GetUserByEmail(email string) (*schema.User, error)
	InsertUser(u *schema.User) (*schema.User, error)
	GetAll(condition *schema.User, opt PaginationOpt) (*[]schema.User, error)
	UpdateById(userId string, data map[string]interface{}) (*schema.User, error)
	DeleteById(userId string) error
}

type DBUserStore struct {
	client *gorm.DB
}

func NewUserStore(client *gorm.DB) *DBUserStore {
	return &DBUserStore{
		client: client,
	}
}

func (s *DBUserStore) Pagination(condition *schema.User, opt PaginationOpt) (*PaginationValue, error) {
	var count int64
	err := s.client.Model(&schema.User{}).Where(&condition).Count(&count).Error
	if err != nil {
		return nil, errors.New("unable to count record")
	}
	return &PaginationValue{PageNum: opt.Page, TotalItems: int(count), TotalPages: int(math.Ceil(float64(count) / float64(opt.Limit)))}, nil
}

func (s *DBUserStore) GetAll(condition *schema.User, opt PaginationOpt) (*[]schema.User, error) {
	var c []schema.User
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.User{}).
		Select("id", "name", "email", "blocked").
		Where(&condition).
		Limit(opt.Limit).
		Offset(o).
		Order("created_at desc").
		Find(&c).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
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

func (s *DBUserStore) UpdateById(userId string, user map[string]interface{}) (*schema.User, error) {
	var result schema.User
	count := s.client.Model(&schema.User{}).Where("id = ?", userId).Updates(user).Scan(&result).RowsAffected
	if count == 0 {
		return nil, errors.New("unable to update user")
	}
	return &result, nil
}

func (s *DBUserStore) DeleteById(userId string) error {
	count := s.client.Where("id = ?", userId).Unscoped().Delete(&schema.User{}).RowsAffected
	if count == 0 {
		return errors.New("unable to update user")
	}
	return nil
}
