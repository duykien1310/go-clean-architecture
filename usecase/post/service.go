package post

import (
	"auth_service/config"
	"auth_service/entity"
	"fmt"
	"time"
)

type Service struct {
	postRepo PostRepository
	userRepo UserRepository
	caching  Caching
}

func NewService(postRepo PostRepository, userRepo UserRepository, caching Caching) *Service {
	return &Service{
		postRepo: postRepo,
		userRepo: userRepo,
		caching:  caching,
	}
}

func (s Service) CreatePost(post *entity.Post) error {
	post, err := s.postRepo.CreatePost(post)
	if err != nil {
		return err
	}

	followerIds, err := s.userRepo.GetListFolowerIdByUserId(post.UserId)
	if err != nil {
		return err
	}
	if len(followerIds) == 0 {
		return nil
	}

	go func(postId int, createdAt time.Time, followerIds []int) {
		for _, userId := range followerIds {
			if err := s.caching.CachePost(
				userId,
				postId,
				float64(createdAt.Unix()),
				config.GetInt("postAge"),
			); err != nil {
				fmt.Printf("failed to cache post for user %d: %v\n", userId, err)
			}
		}
	}(post.Id, post.CreatedAt, followerIds)

	return nil
}

func (s Service) GenerateNewsfeed(userId int) ([]int, error) {
	listPostId, err := s.caching.GetNewsFeedIds(userId)
	if err != nil {
		return nil, err
	}
	if len(listPostId) > 0 {
		return listPostId, nil
	}

	followingUserIds, err := s.userRepo.GetListFolowingUserIdByUserId(userId)
	if err != nil {
		return nil, err
	}
	if len(followingUserIds) == 0 {
		return nil, nil
	}

	postScores, err := s.postRepo.GetPostsByUserIds(followingUserIds)
	if err != nil {
		return nil, err
	}

	err = s.caching.CacheNewsFeed(userId, postScores, config.GetInt("postAge"))
	if err != nil {
		return nil, err
	}

	postIds := make([]int, 0, len(postScores))
	for id := range postScores {
		postIds = append(postIds, id)
	}

	return postIds, nil
}
