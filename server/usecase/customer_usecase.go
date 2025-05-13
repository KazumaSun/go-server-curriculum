package usecase

import (
	"go-server-curriculum/domain"
	"go-server-curriculum/repository"
)

type CustomerUsecase struct {
    customerRepo *repository.CustomerRepository
    productRepo  *repository.ProductRepository
}

func NewCustomerUsecase(customerRepo *repository.CustomerRepository, productRepo *repository.ProductRepository) *CustomerUsecase {
    return &CustomerUsecase{
        customerRepo: customerRepo,
        productRepo:  productRepo,
    }
}

func (u *CustomerUsecase) GetAllCustomers() ([]domain.Customer, error) {
    return u.customerRepo.GetAllCustomers()
}

func (u *CustomerUsecase) GetCustomerByID(id uint) (*domain.Customer, error) {
    return u.customerRepo.GetCustomerByID(id)
}

func (u *CustomerUsecase) CreateCustomer(customer *domain.Customer) (*domain.Customer, error) {
    if err := u.customerRepo.CreateCustomer(customer); err != nil {
        return nil, err
    }
    return customer, nil
}

func (u *CustomerUsecase) UpdateCustomer(id uint, customer *domain.Customer) (*domain.Customer, error) {
    existingCustomer, err := u.customerRepo.GetCustomerByID(id)
    if err != nil {
        return nil, err
    }

    existingCustomer.Name = customer.Name
    existingCustomer.Seat = customer.Seat

    if err := u.customerRepo.UpdateCustomer(existingCustomer); err != nil {
        return nil, err
    }
    return existingCustomer, nil
}

func (u *CustomerUsecase) DeleteCustomer(id uint) error {
    return u.customerRepo.DeleteCustomer(id)
}

func (u *CustomerUsecase) GetCustomerTotal(id uint) (map[string]interface{}, error) {
    customer, err := u.customerRepo.GetCustomerWithOrders(id)
    if err != nil {
        return nil, err
    }

    total := 0
    for _, order := range customer.Orders {
        // ProductRepository を使って Product を取得
        product, err := u.productRepo.GetProductByID(order.ProductID)
        if err != nil {
            return nil, err
        }
        total += order.Quantity * product.Price
    }

    return map[string]interface{}{
        "customer": customer,
        "total":    total,
    }, nil
}