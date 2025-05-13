package usecase

import (
	"go-server-curriculum/domain"
	"go-server-curriculum/repository"
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

// CreateProduct は新しい商品を作成
func (u *ProductUsecase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if err := u.productRepo.CreateProduct(product); err != nil {
			return nil, err
	}
	return product, nil
}

// UpdateProduct は既存の商品を更新
func (u *ProductUsecase) UpdateProduct(id uint, product *domain.Product) (*domain.Product, error) {
	existingProduct, err := u.productRepo.GetProductByID(id)
	if err != nil {
			return nil, err
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price

	if err := u.productRepo.UpdateProduct(existingProduct); err != nil {
			return nil, err
	}
	return existingProduct, nil
}

// DeleteProduct は商品を削除
func (u *ProductUsecase) DeleteProduct(id uint) error {
	return u.productRepo.DeleteProduct(id)
}
