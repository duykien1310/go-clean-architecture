package handler

import (
	userPresenter "auth_service/api/presenter/user"
	"auth_service/config"
	"auth_service/entity"
)

func convertUserDetailToPresenter(data *entity.User) *userPresenter.User {
	listPost := []*userPresenter.Post{}
	for _, post := range data.Post {
		listPost = append(listPost, &userPresenter.Post{
			Id:        post.Id,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Format(config.LAYOUT),
			UpdatedAt: post.UpdatedAt.Format(config.LAYOUT),
		})
	}

	result := &userPresenter.User{
		Id:        data.Id,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Posts:     listPost,

		CreatedAt: data.CreatedAt.Format(config.LAYOUT),
		UpdatedAt: data.UpdatedAt.Format(config.LAYOUT),
	}

	return result
}
