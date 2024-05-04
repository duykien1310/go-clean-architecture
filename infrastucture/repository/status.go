package repository

import (
	"auth_service/entity"

	"gorm.io/gorm"
)

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepository {
	return &StatusRepository{
		db: db,
	}
}

func (r StatusRepository) WithTrx(trxHandle *gorm.DB) StatusRepository {
	if trxHandle == nil {
		return r
	}
	r.db = trxHandle
	return r
}

func (r StatusRepository) GetStatusByCode(code string) (*entity.Status, error) {
	status := entity.Status{}

	result := r.db.Where("code = ?", code).First(&status)
	if result.Error != nil {
		return nil, result.Error
	}

	return &status, nil
}
