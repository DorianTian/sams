package customer_service

import (
	"errors"
	"gorm.io/gorm"
	"sams/app/model"
)

type CustomerService struct {
	DB *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{DB: db}
}

func (s *CustomerService) CreateCustomer(customer *model.Customer) error {
	if err := s.DB.Create(customer).Error; err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) GetCustomers(id string) (*model.Customer, error) {
	var customer model.Customer
	if err := s.DB.First(&customer, "cust_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &customer, nil
}

func (s *CustomerService) GetAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	if err := s.DB.Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}
