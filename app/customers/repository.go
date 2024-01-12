package customers

import (
	"github.com/aditya3232/tes-backend-dbo/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Customers, helper.Pagination, error)
	GetOne(id int) (Customers, error)
	Create(Customers) (Customers, error)
	Update(Customers) (Customers, error)
	Delete(id int) error
	GetByEmail(email string) (Customers, error) // for check unique email
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Customers, helper.Pagination, error) {
	var customers []Customers
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&customers), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return customers, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&customers).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(customers)

	return customers, pagination, nil
}

func (r *repository) GetOne(id int) (Customers, error) {
	var customers Customers

	err := r.db.Where("id = ?", id).First(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (r *repository) Create(customers Customers) (Customers, error) {
	err := r.db.Model(&customers).Create(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (r *repository) Update(customers Customers) (Customers, error) {
	err := r.db.Model(&customers).Where("id = ?", customers.ID).Updates(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}

func (r *repository) Delete(id int) error {
	var customers Customers

	err := r.db.Where("id = ?", id).Delete(&customers).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetByEmail(email string) (Customers, error) {
	var customers Customers

	err := r.db.Where("email = ?", email).First(&customers).Error
	if err != nil {
		return customers, err
	}

	return customers, nil
}
