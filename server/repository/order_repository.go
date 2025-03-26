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
