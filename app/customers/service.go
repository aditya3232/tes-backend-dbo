package customers

import (
	"errors"

	"github.com/aditya3232/tes-backend-dbo/helper"
)

type Service interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Customers, helper.Pagination, error)
	GetOne(input CustomersGetOneByIdInput) (Customers, error)
	Create(input CustomersInput) (Customers, error)
	Update(input CustomersUpdateInput) (Customers, error)
	Delete(input CustomersGetOneByIdInput) error
}

type service struct {
	customersRepository Repository
}

func NewService(customersRepository Repository) *service {
	return &service{customersRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Customers, helper.Pagination, error) {
	customers, pagination, err := s.customersRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return customers, pagination, nil
}

func (s *service) GetOne(input CustomersGetOneByIdInput) (Customers, error) {
	customersID := Customers{ID: input.ID}

	customers, err := s.customersRepository.GetOne(customersID)
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (s *service) Create(input CustomersInput) (Customers, error) {
	email := Customers{Email: input.Email}

	_, err := s.customersRepository.GetByEmail(email)
	if err == nil {
		return Customers{}, errors.New("email must unique")
	}

	customer := Customers{
		UserID:  input.UserID,
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Street:  input.Street,
		ZipCode: input.ZipCode,
		City:    input.City,
		Country: input.Country,
	}

	newCustomers, err := s.customersRepository.Create(customer)
	if err != nil {
		return newCustomers, err
	}

	return newCustomers, nil
}

func (s *service) Update(input CustomersUpdateInput) (Customers, error) {
	customersID := Customers{ID: input.ID}

	_, err := s.customersRepository.GetOne(customersID)
	if err != nil {
		return Customers{}, err
	}

	email := Customers{Email: input.Email}

	_, err = s.customersRepository.GetByEmail(email)
	if err == nil {
		return Customers{}, errors.New("email must unique")
	}

	if input.UserID == nil || *input.UserID == 0 {
		input.UserID = nil
	}

	customer := Customers{
		ID:      input.ID,
		UserID:  input.UserID,
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Street:  input.Street,
		ZipCode: input.ZipCode,
		City:    input.City,
		Country: input.Country,
	}

	newCustomers, err := s.customersRepository.Update(customer)
	if err != nil {
		return newCustomers, err
	}

	return newCustomers, nil

}

func (s *service) Delete(input CustomersGetOneByIdInput) error {
	customersID := Customers{ID: input.ID}

	_, err := s.customersRepository.GetOne(customersID)
	if err != nil {
		return err
	}

	err = s.customersRepository.Delete(customersID)
	if err != nil {
		return err
	}

	return nil
}
