package repository

import (
	"go-server-curriculum/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetAllOrders() ([]domain.Order, error) {
	var orders []domain.Order
	result := r.db.Find(&orders)
	return orders, result.Error
}

// CreateOrder は新しい注文を作成
func (r *OrderRepository) CreateOrder(order *domain.Order) error {
	return r.db.Create(order).Error
}

// GetOrderByID はIDで注文を取得
func (r *OrderRepository) GetOrderByID(id uint) (*domain.Order, error) {
	var order domain.Order
	result := r.db.First(&order, id)
	if result.Error != nil {
			return nil, result.Error
	}
	return &order, nil
}

// UpdateOrder は既存の注文を更新
func (r *OrderRepository) UpdateOrder(order *domain.Order) error {
	return r.db.Save(order).Error
}

// DeleteOrder は注文を削除
func (r *OrderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&domain.Order{}, id).Error
}