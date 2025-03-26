package handler

import (
	"net/http"

	"go-server-curriculum/usecase"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUsecase *usecase.OrderUsecase
}

// NewOrderHandler は OrderHandler を初期化
func NewOrderHandler(orderUsecase *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{orderUsecase: orderUsecase}
}

// GetOrders はすべての注文を取得
func (h *OrderHandler) GetOrders(c echo.Context) error {
	orders, err := h.orderUsecase.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch orders"})
	}
	return c.JSON(http.StatusOK, orders)
}
