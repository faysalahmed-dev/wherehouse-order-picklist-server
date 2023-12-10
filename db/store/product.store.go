package store

import (
	"errors"
	"math"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type ProductStore interface {
	Pagination(condition *schema.Product, opt PaginationOpt) (*PaginationValue, error)
	GetAll(condition *schema.Product, opt PaginationOpt, userId string) (*[]schema.Product, error)
	GetOptions(condition *schema.Product, opt PaginationOpt) (*[]schema.Product, error)
	InsertOne(c *schema.Product) (*schema.Product, error)
	DeleteById(id string) error
	DeleteByUserAndId(userId string, id string) error
	GetByFields(*schema.Product) (*schema.Product, error)
	UpdateOne(condition *schema.Product, data *schema.Product) (*schema.Product, error)
}

type DBProductStore struct {
	client *gorm.DB
}

func NewProductStore(client *gorm.DB) *DBProductStore {
	return &DBProductStore{
		client: client,
	}
}

func (s *DBProductStore) Pagination(condition *schema.Product, opt PaginationOpt) (*PaginationValue, error) {
	var count int64
	err := s.client.Model(&schema.Product{}).Where(&condition).Count(&count).Error
	if err != nil {
		return nil, errors.New("unable to count record")
	}
	return &PaginationValue{PageNum: opt.Page, TotalItems: int(count), TotalPages: int(math.Ceil(float64(count) / float64(opt.Limit)))}, nil
}

func (s *DBProductStore) GetAll(condition *schema.Product, opt PaginationOpt, userId string) (*[]schema.Product, error) {
	var c []schema.Product
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.Product{}).
		Where(condition).
		Select("id", "name", "user_id", "sub_category_id").
		Limit(opt.Limit).
		Offset(o).
		Order("created_at desc").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "type")
		}).
		Preload("SubCategory", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "value")
		}).
		Preload("Orders", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "amount", "unit_type", "product_id").Where("user_id = ?", userId)
		}).
		Find(&c).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}

func (s *DBProductStore) GetOptions(condition *schema.Product, opt PaginationOpt) (*[]schema.Product, error) {
	var c []schema.Product
	o := (opt.Page - 1) * opt.Limit
	err := s.client.Model(&schema.Product{}).
		Where(condition).
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

func (s *DBProductStore) InsertOne(c *schema.Product) (*schema.Product, error) {
	result := s.client.Create(c)
	if result.Error != nil {
		return nil, errors.New("unable to insert record")
	}
	return c, nil
}

func (s *DBProductStore) DeleteById(id string) error {
	count := s.client.Where("id = ?", id).Unscoped().Delete(&schema.Product{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (s *DBProductStore) DeleteByUserAndId(userId string, id string) error {
	count := s.client.Where("user_id = ? AND id = ?", userId, id).Unscoped().Delete(&schema.Product{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}

func (s *DBProductStore) GetByFields(c *schema.Product) (*schema.Product, error) {
	var result *schema.Product
	err := s.client.Where(c).First(&result).Error
	return result, err
}

func (s *DBProductStore) UpdateOne(condition *schema.Product, data *schema.Product) (*schema.Product, error) {
	var result schema.Product
	r := s.client.Model(&schema.Product{}).Where(&condition).Updates(data).Scan(&result)
	if r.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return &result, nil
}
