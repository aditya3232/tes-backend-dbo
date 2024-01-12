package orders

import "github.com/aditya3232/tes-backend-dbo/helper"

type Service interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Orders, helper.Pagination, error)
	GetOne(input OrdersGetOneByIdInput) (Orders, error)
	Create(input OrdersInput) (Orders, error)
	Update(input OrdersUpdateInput) (Orders, error)
	Delete(input OrdersGetOneByIdInput) error
}

type service struct {
	ordersRepository Repository
}

func NewService(ordersRepository Repository) *service {
	return &service{ordersRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Orders, helper.Pagination, error) {
	orders, pagination, err := s.ordersRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return orders, pagination, nil
}

func (s *service) GetOne(input OrdersGetOneByIdInput) (Orders, error) {
	orders, err := s.ordersRepository.GetOne(input.ID)
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *service) Create(input OrdersInput) (Orders, error) {
	orders := Orders{
		CustomerID:  input.CustomerID,
		TotalAmount: input.TotalAmount,
		Status:      input.Status,
		PaymentType: input.PaymentType,
	}

	newOrders, err := s.ordersRepository.Create(orders)
	if err != nil {
		return newOrders, err
	}

	return newOrders, nil
}

func (s *service) Update(input OrdersUpdateInput) (Orders, error) {
	_, err := s.ordersRepository.GetOne(input.ID)
	if err != nil {
		return Orders{}, err
	}

	if input.CustomerID == nil || *input.CustomerID == 0 {
		input.CustomerID = nil
	}

	orders := Orders{
		ID:          input.ID,
		CustomerID:  input.CustomerID,
		TotalAmount: input.TotalAmount,
		Status:      input.Status,
		PaymentType: input.PaymentType,
	}

	newOrders, err := s.ordersRepository.Update(orders)
	if err != nil {
		return newOrders, err
	}

	return newOrders, nil
}

func (s *service) Delete(input OrdersGetOneByIdInput) error {
	_, err := s.ordersRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.ordersRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
