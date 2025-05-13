package handler

import (
	"net/http"
	"strconv"

	"go-server-curriculum/domain"
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

// CreateOrder は新しい注文を作成
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var order domain.Order
	if err := c.Bind(&order); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	createdOrder, err := h.orderUsecase.CreateOrder(&order)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}
	return c.JSON(http.StatusCreated, createdOrder)
}

// GetOrder はIDで注文を取得
func (h *OrderHandler) GetOrder(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	order, err := h.orderUsecase.GetOrderByID(uint(id))
	if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	return c.JSON(http.StatusOK, order)
}

// UpdateOrder は既存の注文を更新
func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	var order domain.Order
	if err := c.Bind(&order); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedOrder, err := h.orderUsecase.UpdateOrder(uint(id), &order)
	if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order"})
	}
	return c.JSON(http.StatusOK, updatedOrder)
}

// DeleteOrder は注文を削除
func (h *OrderHandler) DeleteOrder(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order ID"})
	}

	if err := h.orderUsecase.DeleteOrder(uint(id)); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete order"})
	}
	return c.JSON(http.StatusNoContent, nil)
}
