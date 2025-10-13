package post

import "auth_service/entity"

type PostRepository interface {
	CreatePost(post *entity.Post) error
}

type UseCase interface {
	CreatePost(post *entity.Post) error
}
