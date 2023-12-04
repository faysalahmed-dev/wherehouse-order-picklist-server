package store

import (
	"errors"
	"math"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type ProductStore interface {
	Pagination(limit int, condition *schema.Product) (pageNum int, err error)
	GetAll(page int, limit int, condition *schema.Product) (*[]schema.Product, error)
	GetOptions(page int, limit int, condition *schema.Product) (*[]schema.Product, error)
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

func (s *DBProductStore) Pagination(limit int, condition *schema.Product) (int, error) {
	var count int64
	err := s.client.Model(&schema.Product{}).Where(condition).Count(&count).Error
	if err != nil {
		return 0, errors.New("unable to count record")
	}
	return int(math.Ceil(float64(count) / float64(limit))), nil
}

func (s *DBProductStore) GetAll(page int, limit int, condition *schema.Product) (*[]schema.Product, error) {
	var c []schema.Product
	o := (page - 1) * limit
	err := s.client.Model(&schema.Product{}).
		Where(condition).
		Select("id", "name", "value", "description", "user_id", "sub_category_id").
		Limit(limit).
		Offset(o).
		Order("created_at desc").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name")
		}).
		Preload("SubCategory", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "value")
		}).
		Find(&c).Error
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &c, nil
}

func (s *DBProductStore) GetOptions(page int, limit int, condition *schema.Product) (*[]schema.Product, error) {
	var c []schema.Product
	o := (page - 1) * limit
	err := s.client.Model(&schema.Product{}).
		Where(condition).
		Select("id", "name", "value").
		Limit(limit).
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
	return s.client.Unscoped().Delete(&schema.Product{}, id).Error
}

func (s *DBProductStore) DeleteByUserAndId(userId string, id string) error {
	return s.client.Where("user_id = ? AND id = ?", userId, id).Unscoped().Delete(&schema.Product{}).Error
}

func (s *DBProductStore) GetByFields(c *schema.Product) (*schema.Product, error) {
	var result *schema.Product
	err := s.client.Where(c).First(&result).Error
	return result, err
}

func (s *DBProductStore) UpdateOne(condition *schema.Product, data *schema.Product) (*schema.Product, error) {
	var result *schema.Product
	r := s.client.Model(&schema.Product{}).Where(&condition).Updates(data)

	return result, r.Error
}
