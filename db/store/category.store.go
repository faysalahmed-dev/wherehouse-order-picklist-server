package store

import (
	"errors"
	"math"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type CategoryStore interface {
	Pagination(c *schema.Category, opt PaginationOpt) (*PaginationValue, error)
	GetCategories(opt PaginationOpt) (*[]schema.Category, error)
	GetCategoryOptions(opt PaginationOpt) (*[]schema.Category, error)
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

func (s *DBCategoryStore) Pagination(c *schema.Category, opt PaginationOpt) (*PaginationValue, error) {
	var count int64
	err := s.client.Model(&schema.Category{}).Where(&c).Count(&count).Error
	if err != nil {
		return nil, errors.New("unable to count record")
	}
	return &PaginationValue{PageNum: opt.Page, TotalItems: int(count), TotalPages: int(math.Ceil(float64(count) / float64(opt.Limit)))}, nil
}
func (s *DBCategoryStore) GetCategories(opt PaginationOpt) (*[]schema.Category, error) {
	var c []schema.Category
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.Category{}).
		Select("id", "name", "value", "user_id").
		Limit(opt.Limit).
		Offset(o).
		Order("created_at desc").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "type")
		}).
		Find(&c).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}
func (s *DBCategoryStore) GetCategoryOptions(opt PaginationOpt) (*[]schema.Category, error) {
	var c []schema.Category
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.Category{}).
		Select("id", "name", "value").
		Limit(opt.Limit).
		Offset(o).
		Order("created_at desc").
		Find(&c).Error
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
	count := s.client.Where("id = ?", id).Unscoped().Delete(&schema.Category{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (s *DBCategoryStore) DeleteByUserAndId(userId string, id string) error {
	count := s.client.Where("user_id = ? AND id = ?", userId, id).Unscoped().Delete(&schema.Category{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (s *DBCategoryStore) GetByFields(c *schema.Category) (*schema.Category, error) {
	var result *schema.Category
	err := s.client.Where(c).First(&result).Error
	return result, err
}

func (s *DBCategoryStore) UpdateById(id string, c *schema.Category) (*schema.Category, error) {
	var result *schema.Category
	r := s.client.Model(&schema.Category{}).Where("id = ?", id).Updates(c).Scan(&result)
	if r.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}
