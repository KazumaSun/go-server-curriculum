package repository

import (
	"go-server-curriculum/server/domain"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository は ProductRepository を初期化
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// GetAllProducts はすべての商品を取得
func (r *ProductRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	result := r.db.Find(&products)
	return products, result.Error
}

// GetProductByID はIDから商品を取得
func (r *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
