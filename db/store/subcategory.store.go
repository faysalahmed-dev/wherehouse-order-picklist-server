package store

import (
	"errors"
	"fmt"
	"math"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type SubCategoryStore interface {
	Pagination(condition *schema.SubCategory, opt PaginationOpt) (*PaginationValue, error)
	GetAll(condition *schema.SubCategory, opt PaginationOpt) (*[]schema.SubCategory, error)
	GetOptions(condition *schema.SubCategory, opt PaginationOpt) (*[]schema.SubCategory, error)
	InsertOne(c *schema.SubCategory) (*schema.SubCategory, error)
	DeleteById(id string) error
	DeleteByUserAndId(userId string, id string) error
	GetByFields(*schema.SubCategory) (*schema.SubCategory, error)
	UpdateOne(condition *schema.SubCategory, data *schema.SubCategory) (*schema.SubCategory, error)
}

type DBSubCategoryStore struct {
	client *gorm.DB
}

func NewSubCategoryStore(client *gorm.DB) *DBSubCategoryStore {
	return &DBSubCategoryStore{
		client: client,
	}
}

func (s *DBSubCategoryStore) Pagination(condition *schema.SubCategory, opt PaginationOpt) (*PaginationValue, error) {
	var count int64
	err := s.client.Model(&schema.SubCategory{}).Where(&condition).Count(&count).Error
	if err != nil {
		return nil, errors.New("unable to count record")
	}
	return &PaginationValue{PageNum: opt.Page, TotalItems: int(count), TotalPages: int(math.Ceil(float64(count) / float64(opt.Limit)))}, nil
}

func (s *DBSubCategoryStore) GetAll(condition *schema.SubCategory, opt PaginationOpt) (*[]schema.SubCategory, error) {
	var c []schema.SubCategory
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.SubCategory{}).
		Where(condition).
		Select("id", "name", "value", "description", "user_id", "category_id").
		Limit(opt.Limit).
		Offset(o).
		Order("created_at desc").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "type")
		}).
		Preload("Category", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "value")
		}).
		Find(&c).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}

func (s *DBSubCategoryStore) GetOptions(condition *schema.SubCategory, opt PaginationOpt) (*[]schema.SubCategory, error) {
	var c []schema.SubCategory
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.Category{}).
		Where(condition).
		Select("id", "name", "value").
		Limit(opt.Limit).
		Offset(o).
		Order("created_at desc").
		Find(&c).Error
	fmt.Println(err)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}

func (s *DBSubCategoryStore) InsertOne(c *schema.SubCategory) (*schema.SubCategory, error) {
	result := s.client.Create(c)
	if result.Error != nil {
		return nil, errors.New("unable to insert record")
	}
	return c, nil
}

func (s *DBSubCategoryStore) DeleteById(id string) error {
	count := s.client.Where("id = ?", id).Unscoped().Delete(&schema.SubCategory{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (s *DBSubCategoryStore) DeleteByUserAndId(userId string, id string) error {
	count := s.client.Where("user_id = ? AND id = ?", userId, id).Unscoped().Delete(&schema.SubCategory{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (s *DBSubCategoryStore) GetByFields(c *schema.SubCategory) (*schema.SubCategory, error) {
	var result *schema.SubCategory
	err := s.client.Where(c).First(&result).Error
	return result, err
}

func (s *DBSubCategoryStore) UpdateOne(condition *schema.SubCategory, data *schema.SubCategory) (*schema.SubCategory, error) {
	var result *schema.SubCategory
	r := s.client.Model(&schema.SubCategory{}).Where(&condition).Updates(data).Scan(&result)
	if r.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return result, nil
}
