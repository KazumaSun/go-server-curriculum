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
