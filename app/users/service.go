package users

import (
	"errors"

	"github.com/aditya3232/tes-backend-dbo/helper"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Users, helper.Pagination, error)
	GetOne(input UsersGetOneByIdInput) (Users, error)
	Create(input UsersInput) (Users, error)
	Update(input UsersUpdateInput) (Users, error)
	Delete(input UsersGetOneByIdInput) error
}

type service struct {
	userRepository Repository
}

func NewService(userRepository Repository) *service {
	return &service{userRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Users, helper.Pagination, error) {
	users, pagination, err := s.userRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return users, pagination, nil
}

func (s *service) GetOne(input UsersGetOneByIdInput) (Users, error) {
	user, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Create(input UsersInput) (Users, error) {
	_, err := s.userRepository.GetUsername(input.Username)
	if err == nil {
		return Users{}, errors.New("username must unique")
	}

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return Users{}, err
		}

		input.Password = string(password)
	}

	user := Users{
		Username: input.Username,
		Password: input.Password,
	}

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Update(input UsersUpdateInput) (Users, error) {
	_, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return Users{}, err
	}

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return Users{}, err
		}

		input.Password = string(password)
	}

	user := Users{
		ID:            input.ID,
		Password:      input.Password,
		RememberToken: input.RememberToken, // dari login
	}

	newUser, err := s.userRepository.Update(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Delete(input UsersGetOneByIdInput) error {
	_, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
