package repository

import (
	"go-server-curriculum/domain"

	"gorm.io/gorm"
)

type CustomerRepository struct {
    db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
    return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAllCustomers() ([]domain.Customer, error) {
    var customers []domain.Customer
    err := r.db.Find(&customers).Error
    return customers, err
}

func (r *CustomerRepository) GetCustomerByID(id uint) (*domain.Customer, error) {
    var customer domain.Customer
    err := r.db.First(&customer, id).Error
    return &customer, err
}

func (r *CustomerRepository) CreateCustomer(customer *domain.Customer) error {
    return r.db.Create(customer).Error
}

func (r *CustomerRepository) UpdateCustomer(customer *domain.Customer) error {
    return r.db.Save(customer).Error
}

func (r *CustomerRepository) DeleteCustomer(id uint) error {
    return r.db.Delete(&domain.Customer{}, id).Error
}

func (r *CustomerRepository) GetCustomerWithOrders(id uint) (*domain.Customer, error) {
    var customer domain.Customer
    err := r.db.Preload("Orders").First(&customer, id).Error
    if err != nil {
        return nil, err
    }
    return &customer, nil
}