package logindata

import "github.com/aditya3232/tes-backend-dbo/app/users"

type Service interface {
	GetLoginData(userID int) (users.Users, error)
}

type service struct {
	usersRepository users.Repository
}

func NewService(usersRepository users.Repository) *service {
	return &service{usersRepository}
}

func (s *service) GetLoginData(userID int) (users.Users, error) {
	user, err := s.usersRepository.GetOne(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}
