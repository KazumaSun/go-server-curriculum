package usecase

import (
	"go-server-curriculum/server/domain"
	"go-server-curriculum/server/repository"
)

type ProductUsecase struct {
	productRepo *repository.ProductRepository
}

// NewProductUsecase は ProductUsecase を初期化
func NewProductUsecase(productRepo *repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo: productRepo}
}

// GetAllProducts はすべての商品を取得
func (u *ProductUsecase) GetAllProducts() ([]domain.Product, error) {
	return u.productRepo.GetAllProducts()
}

// GetProductByID はIDで商品を取得
func (u *ProductUsecase) GetProductByID(id uint) (*domain.Product, error) {
	return u.productRepo.GetProductByID(id)
}
