package store

import (
	"errors"
	"fmt"
	"math"

	"github.com/faysalahmed-dev/wherehouse-order-picklist/db/schema"
	"gorm.io/gorm"
)

type OrderStore interface {
	Pagination(c *schema.Order, opt PaginationOpt) (*PaginationValue, error)
	InsertOne(c *schema.Order) (*schema.Order, error)
	GroupOrderByUser(userId string, opt PaginationOpt) (*[]schema.GroupOrderByUser, error)
	OrdersByUserId(userId string, opt PaginationOpt) (*[]schema.Order, error)
	HasOrderUsers(opt PaginationOpt) (*[]schema.OrderUserOptions, error)
	HasOrder(pId string, uId string) bool
	DeleteByFields(condition *schema.Order) error
	UpdateByFields(condition *schema.Order, data *schema.Order) (*schema.Order, error)
}

type DBOrderStore struct {
	client *gorm.DB
}

func NewOrderStore(client *gorm.DB) *DBOrderStore {
	return &DBOrderStore{
		client: client,
	}
}
func (s *DBOrderStore) Pagination(c *schema.Order, opt PaginationOpt) (*PaginationValue, error) {
	var count int64
	err := s.client.Model(&schema.Order{}).Where(&c).Count(&count).Error
	if err != nil {
		return nil, errors.New("unable to count record")
	}
	return &PaginationValue{PageNum: opt.Page, TotalItems: int(count), TotalPages: int(math.Ceil(float64(count) / float64(opt.Limit)))}, nil
}

func (s *DBOrderStore) InsertOne(o *schema.Order) (*schema.Order, error) {
	return o, s.client.Create(o).Error
}

func (s *DBOrderStore) GroupOrderByUser(userId string, opt PaginationOpt) (*[]schema.GroupOrderByUser, error) {
	o := (opt.Page - 1) * opt.Limit
	var results []schema.GroupOrderByUser
	var filter = ""
	if len(userId) > 0 {
		filter = fmt.Sprintf("user_id = '%v'", userId)
	}
	if err := s.client.Table("orders o").Select("count(o.id) as total, SUM(CASE WHEN o.status = 'DONE' THEN 1 ELSE 0 END) as picked, user_id, u.name, u.email, max(o.updated_at) as last_submission").Where(filter).Limit(opt.Limit).Offset(o).Joins("JOIN users u ON u.id = o.user_id").Group("u.name, u.email, o.user_id").Order("last_submission desc").Scan(&results).Error; err != nil {
		return nil, err
	}
	return &results, nil
}

func (s *DBOrderStore) OrdersByUserId(userId string, opt PaginationOpt) (*[]schema.Order, error) {
	o := (opt.Page - 1) * opt.Limit
	var results []schema.Order

	if err := s.client.Model(&schema.Order{}).Where("user_id = ? ", userId).Preload("Product", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "sub_category_id").Preload("SubCategory", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "name", "value", "category_id").Preload("Category", func(db *gorm.DB) *gorm.DB {
				return db.Select("id", "name", "value")
			})
		})
	}).Limit(opt.Limit).Offset(o).Order("updated_at desc").Find(&results).Error; err != nil {
		return nil, err
	}

	return &results, nil
}

func (s *DBOrderStore) HasOrderUsers(opt PaginationOpt) (*[]schema.OrderUserOptions, error) {
	var results []schema.OrderUserOptions

	s.client.Table("orders o").Select("user_id as id, u.name").Joins("JOIN users u ON u.id = o.user_id").Group("u.name, o.user_id").Limit(50).Scan(&results)
	return &results, nil
}

func (s *DBOrderStore) HasOrder(pId string, uId string) bool {
	var exists bool

	if err := s.client.Model(&schema.Order{}).Where(&schema.Order{UserId: uId, ProductId: pId}).Find(&exists).Error; err != nil || !exists {
		return false
	}
	return exists
}

func (s *DBOrderStore) DeleteByFields(condition *schema.Order) error {
	count := s.client.Where(condition).Unscoped().Delete(&schema.Order{}).RowsAffected
	if count == 0 {
		return errors.New("item not found")
	}
	return nil
}
func (s *DBOrderStore) UpdateByFields(condition *schema.Order, data *schema.Order) (*schema.Order, error) {
	var result *schema.Order
	count := s.client.Model(&schema.Order{}).Where(&condition).Updates(data).Scan(&result).RowsAffected
	if count == 0 {
		return nil, errors.New("order not found")
	}
	return result, nil
}
