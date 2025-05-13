package usecase

import (
	"go-server-curriculum/domain"
	"go-server-curriculum/repository"
)

type OrderUsecase struct {
	orderRepo *repository.OrderRepository
}

func NewOrderUsecase(orderRepo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: orderRepo}
}

func (u *OrderUsecase) GetAllOrders() ([]domain.Order, error) {
	return u.orderRepo.GetAllOrders()
}

// CreateOrder は新しい注文を作成
func (u *OrderUsecase) CreateOrder(order *domain.Order) (*domain.Order, error) {
	if err := u.orderRepo.CreateOrder(order); err != nil {
			return nil, err
	}
	return order, nil
}

// GetOrderByID はIDで注文を取得
func (u *OrderUsecase) GetOrderByID(id uint) (*domain.Order, error) {
	return u.orderRepo.GetOrderByID(id)
}

// UpdateOrder は既存の注文を更新
func (u *OrderUsecase) UpdateOrder(id uint, order *domain.Order) (*domain.Order, error) {
	existingOrder, err := u.orderRepo.GetOrderByID(id)
	if err != nil {
			return nil, err
	}

	existingOrder.ProductID = order.ProductID
	existingOrder.Quantity = order.Quantity

	if err := u.orderRepo.UpdateOrder(existingOrder); err != nil {
			return nil, err
	}
	return existingOrder, nil
}

// DeleteOrder は注文を削除
func (u *OrderUsecase) DeleteOrder(id uint) error {
	return u.orderRepo.DeleteOrder(id)
}