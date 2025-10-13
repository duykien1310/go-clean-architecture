package post

import "auth_service/entity"

type Service struct {
	postRepo PostRepository
}

func NewService(postRepo PostRepository) *Service {
	return &Service{
		postRepo: postRepo,
	}
}

func (s Service) CreatePost(post *entity.Post) error {
	return s.postRepo.CreatePost(post)
}
