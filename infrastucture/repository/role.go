package repository

import (
	"auth_service/entity"

	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r RoleRepository) WithTrx(trxHandle *gorm.DB) RoleRepository {
	if trxHandle == nil {
		return r
	}
	r.db = trxHandle
	return r
}

func (r RoleRepository) GetRoleByCode(code string) (*entity.Role, error) {
	role := entity.Role{}

	result := r.db.Where("code = ?", code).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}
