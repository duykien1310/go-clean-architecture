package repository

import (
	"auth_service/entity"
	"time"

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

func (r PostRepository) CreatePost(post *entity.Post) (*entity.Post, error) {
	result := r.db.Create(post)
	if result.Error != nil {
		return nil, result.Error
	}

	return post, nil
}

func (r PostRepository) GetPostsByUserIds(listUserId []int) (map[int]time.Time, error) {
	posts := []*entity.Post{}
	err := r.db.Model(&entity.Post{}).
		Select("id, created_at").
		Where("user_id IN ?", listUserId).
		Order("created_at DESC").
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	postScores := make(map[int]time.Time)
	for _, p := range posts {
		postScores[p.Id] = p.CreatedAt
	}
	return postScores, nil
}
