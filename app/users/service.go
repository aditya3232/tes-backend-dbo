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
	userID := Users{ID: input.ID}

	user, err := s.userRepository.GetOne(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Create(input UsersInput) (Users, error) {
	username := Users{Username: input.Username}

	_, err := s.userRepository.GetUsername(username)
	if err == nil {
		return Users{}, errors.New("username must unique")
	}

	if input.RoleID == nil || *input.RoleID == 0 {
		roleID := 34
		input.RoleID = &roleID
	}

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return Users{}, err
		}

		input.Password = string(password)
	}

	user := Users{
		RoleID:   input.RoleID,
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
	userID := Users{ID: input.ID}

	_, err := s.userRepository.GetOne(userID)
	if err != nil {
		return Users{}, err
	}

	// if roleID == nil || *roleID == 0, then dont update roleID to default int value 0
	if input.RoleID == nil || *input.RoleID == 0 {
		input.RoleID = nil
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
		RoleID:        input.RoleID,
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
	userID := Users{ID: input.ID}

	_, err := s.userRepository.GetOne(userID)
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
