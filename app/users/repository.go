package users

import (
	"github.com/aditya3232/tes-backend-dbo/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Users, helper.Pagination, error)
	GetOne(Users) (Users, error)
	Create(Users) (Users, error)
	Update(Users) (Users, error)
	Delete(Users) error
	GetUsername(Users) (Users, error) // for check unique username
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Users, helper.Pagination, error) {
	var users []Users
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&users), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return users, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&users).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(users)

	return users, pagination, nil
}

func (r *repository) GetOne(user Users) (Users, error) {
	err := r.db.Where("id = ?", user.ID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user Users) (Users, error) {
	err := r.db.Model(&user).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user Users) (Users, error) {
	err := r.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(user Users) error {
	err := r.db.Where("id = ?", user.ID).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUsername(user Users) (Users, error) {
	err := r.db.Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
