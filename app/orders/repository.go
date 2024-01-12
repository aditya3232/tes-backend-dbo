package orders

import (
	"github.com/aditya3232/tes-backend-dbo/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Orders, helper.Pagination, error)
	GetOne(id int) (Orders, error)
	Create(Orders) (Orders, error)
	Update(Orders) (Orders, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Orders, helper.Pagination, error) {
	var orders []Orders
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&orders), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return orders, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&orders).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(orders)

	return orders, pagination, nil
}

func (r *repository) GetOne(id int) (Orders, error) {
	var orders Orders

	err := r.db.Where("id = ?", id).First(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *repository) Create(orders Orders) (Orders, error) {
	err := r.db.Model(&orders).Create(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *repository) Update(orders Orders) (Orders, error) {
	err := r.db.Model(&orders).Where("id = ?", orders.ID).Updates(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *repository) Delete(id int) error {
	var orders Orders

	err := r.db.Where("id = ?", id).Delete(&orders).Error
	if err != nil {
		return err
	}

	return nil
}
