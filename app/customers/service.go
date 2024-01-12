package customers

import (
	"errors"

	"github.com/aditya3232/tes-backend-dbo/app/users"
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
	userRepository      users.Repository // for checking user_id
}

func NewService(customersRepository Repository, userRepository users.Repository) *service {
	return &service{customersRepository, userRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Customers, helper.Pagination, error) {
	customers, pagination, err := s.customersRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return customers, pagination, nil
}

func (s *service) GetOne(input CustomersGetOneByIdInput) (Customers, error) {
	customers, err := s.customersRepository.GetOne(input.ID)
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (s *service) Create(input CustomersInput) (Customers, error) {
	_, err := s.customersRepository.GetByEmail(input.Email)
	if err == nil {
		return Customers{}, errors.New("email must unique")
	}

	_, err = s.userRepository.GetOne(*input.UserID)
	if err == nil {
		return Customers{}, errors.New("user must unique")
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
	_, err := s.customersRepository.GetOne(input.ID)
	if err != nil {
		return Customers{}, err
	}

	_, err = s.customersRepository.GetByEmail(input.Email)
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
	_, err := s.customersRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.customersRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
