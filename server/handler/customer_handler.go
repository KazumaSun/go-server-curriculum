package handler

import (
	"net/http"
	"strconv"

	"go-server-curriculum/domain"
	"go-server-curriculum/usecase"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	customerUsecase *usecase.CustomerUsecase
}

func NewCustomerHandler(customerUsecase *usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{customerUsecase: customerUsecase}
}

// GetCustomers はすべての顧客を取得
func (h *CustomerHandler) GetCustomers(c echo.Context) error {
	customers, err := h.customerUsecase.GetAllCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch customers"})
	}
	return c.JSON(http.StatusOK, customers)
}

// GetCustomer はIDで顧客を取得
func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	customer, err := h.customerUsecase.GetCustomerByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Customer not found"})
	}
	return c.JSON(http.StatusOK, customer)
}

// CreateCustomer は新しい顧客を作成
func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	createdCustomer, err := h.customerUsecase.CreateCustomer(&customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create customer"})
	}
	return c.JSON(http.StatusCreated, createdCustomer)
}

// UpdateCustomer は既存の顧客を更新
func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	var customer domain.Customer
	if err := c.Bind(&customer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	updatedCustomer, err := h.customerUsecase.UpdateCustomer(uint(id), &customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update customer"})
	}
	return c.JSON(http.StatusOK, updatedCustomer)
}

// DeleteCustomer は顧客を削除
func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	if err := h.customerUsecase.DeleteCustomer(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete customer"})
	}
	return c.JSON(http.StatusNoContent, nil)
}

// GetCustomerTotal は顧客の合計を取得
func (h *CustomerHandler) GetCustomerTotal(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid customer ID"})
	}

	result, err := h.customerUsecase.GetCustomerTotal(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch customer total"})
	}

	return c.JSON(http.StatusOK, result)
}