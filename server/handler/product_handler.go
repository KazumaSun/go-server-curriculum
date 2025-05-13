package handler

import (
	"net/http"
	"strconv"

	"go-server-curriculum/domain"
	"go-server-curriculum/usecase"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

// NewProductHandler は ProductHandler を初期化
func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

// GetProducts は商品一覧を取得
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch products"})
	}
	return c.JSON(http.StatusOK, products)
}

// GetProduct はIDで商品を取得
func (h *ProductHandler) GetProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	product, err := h.productUsecase.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

// CreateProduct は新しい商品を作成
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var product domain.Product
	if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	createdProduct, err := h.productUsecase.CreateProduct(&product)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create product"})
	}
	return c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct は既存の商品を更新
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var product domain.Product
	if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedProduct, err := h.productUsecase.UpdateProduct(uint(id), &product)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}
	return c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct は商品を削除
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	if err := h.productUsecase.DeleteProduct(uint(id)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}
	return c.JSON(http.StatusNoContent, nil)
}