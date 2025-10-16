package repository

import (
	"auth_service/entity"

	"gorm.io/gorm"
)

type FollowingRepository struct {
	db *gorm.DB
}

func NewFollowingRepository(db *gorm.DB) *FollowingRepository {
	return &FollowingRepository{
		db: db,
	}
}

func (r FollowingRepository) FollowUser(following *entity.Following) error {
	result := r.db.Create(following)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
