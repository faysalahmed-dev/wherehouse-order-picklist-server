package store

import (
	"errors"
	"fmt"
	"math"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type CategoryStore interface {
	Pagination(limit int) (pageNum int, err error)
	GetCategories(page int, limit int) (*[]schema.Category, error)
	GetCategoryOptions(page int, limit int) (*[]schema.Category, error)
	InsertCategory(c *schema.Category) (*schema.Category, error)
	DeleteById(id string) error
	DeleteByUserAndId(userId string, id string) error
	GetByFields(*schema.Category) (*schema.Category, error)
	UpdateById(id string, c *schema.Category) (*schema.Category, error)
}

type DBCategoryStore struct {
	client *gorm.DB
}

func NewCategoryStore(client *gorm.DB) *DBCategoryStore {
	return &DBCategoryStore{
		client: client,
	}
}

func (s *DBCategoryStore) Pagination(limit int) (int, error) {
	var count int64
	err := s.client.Model(&schema.Category{}).Count(&count).Error
	fmt.Println(err)
	if err != nil {
		return 0, errors.New("unable to count record")
	}
	return int(math.Ceil(float64(count) / float64(limit))), nil
}
func (s *DBCategoryStore) GetCategories(page int, limit int) (*[]schema.Category, error) {
	var c []schema.Category
	o := (page - 1) * limit
	err := s.client.Model(&schema.Category{}).
		Select("id", "name", "value", "user", "user_id").
		Limit(limit).
		Offset(o).
		Order("created_at desc").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Find(&c).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}

func (s *DBCategoryStore) GetCategoryOptions(page int, limit int) (*[]schema.Category, error) {
	var c []schema.Category
	o := (page - 1) * limit
	err := s.client.Model(&schema.Category{}).
		Select("id", "name", "value").
		Limit(limit).
		Offset(o).
		Order("created_at desc").
		Find(&c).Error
	fmt.Println(err)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}

func (s *DBCategoryStore) InsertCategory(c *schema.Category) (*schema.Category, error) {
	result := s.client.Create(c)
	if result.Error != nil {
		return nil, errors.New("unable to insert user")
	}
	return c, nil
}

func (s *DBCategoryStore) DeleteById(id string) error {
	return s.client.Unscoped().Delete(&schema.Category{}, id).Error
}

func (s *DBCategoryStore) DeleteByUserAndId(userId string, id string) error {
	return s.client.Where("user_id = ? AND id = ?", userId, id).Unscoped().Delete(&schema.Category{}).Error
}

func (s *DBCategoryStore) GetByFields(c *schema.Category) (*schema.Category, error) {
	var result *schema.Category
	err := s.client.Where(c).First(&result).Error
	return result, err
}

func (s *DBCategoryStore) UpdateById(id string, c *schema.Category) (*schema.Category, error) {
	var result *schema.Category
	r := s.client.Model(&schema.Category{}).Where("id = ?", id).Updates(c)

	return result, r.Error
}
