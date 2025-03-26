package handler

import (
	"net/http"
	"strconv"

	"go-server-curriculum/usecase"

	"github.com/gorilla/mux"
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
	idParam := mux.Vars(c.Request())["id"]
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
