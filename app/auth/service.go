package auth

import (
	"time"

	users_app "github.com/aditya3232/tes-backend-dbo/app/users"
	"github.com/aditya3232/tes-backend-dbo/library/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input LoginInput) (users_app.Users, error)
	Logout(token string) error
}

type service struct {
	usersRepository users_app.Repository
}

func NewService(usersRepository users_app.Repository) *service {
	return &service{usersRepository}
}

func (s *service) Login(input LoginInput) (users_app.Users, error) {
	var entityUsers users_app.Users

	user, err := s.usersRepository.GetUsername(input.Username)
	if err != nil {
		return entityUsers, err
	}

	if user.ID == 0 {
		return entityUsers, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return entityUsers, err
	}

	token, err := jwt.GenerateToken(user.ID, 30)
	if err != nil {
		return entityUsers, err
	}

	now := time.Now()

	entityUsers = users_app.Users{
		ID:            user.ID,
		RememberToken: token,
		UpdatedAt:     &now,
	}

	loginUser, err := s.usersRepository.Update(entityUsers)
	if err != nil {
		return loginUser, err
	}

	return loginUser, nil
}

func (s *service) Logout(token string) error {
	userID, err := jwt.GetUserIDFromToken(token)
	if err != nil {
		return err
	}

	users, err := s.usersRepository.GetOne(userID)
	if err != nil {
		return err
	}

	now := time.Now()

	entityUsers := users_app.Users{
		ID:            users.ID,
		RememberToken: " ",
		UpdatedAt:     &now,
	}

	_, err = s.usersRepository.Update(entityUsers)
	if err != nil {
		return err
	}

	return nil
}
