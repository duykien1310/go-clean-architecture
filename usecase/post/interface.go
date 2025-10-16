package post

import (
	"auth_service/entity"
	"time"
)

type PostRepository interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
	GetPostsByUserIds(listUserId []int) (map[int]time.Time, error)
}

type UserRepository interface {
	GetListFolowingUserIdByUserId(userId int) ([]int, error)
	GetListFolowerIdByUserId(userId int) ([]int, error)
}

type Caching interface {
	CacheNewsFeed(userId int, postScores map[int]time.Time, expiresAt int) error
	GetNewsFeedIds(userId int) ([]int, error)
	CachePost(userId int, postId int, scores float64, expiresAt int) error
}

type UseCase interface {
	CreatePost(post *entity.Post) error
	GenerateNewsfeed(userId int) ([]int, error)
}
