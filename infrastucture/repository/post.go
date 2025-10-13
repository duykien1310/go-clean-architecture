package repository

import (
	"auth_service/entity"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r PostRepository) CreatePost(post *entity.Post) error {
	result := r.db.Create(post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
