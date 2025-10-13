package presenter

import (
	"auth_service/api/presenter"
)

type UserResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Results []*User `json:"results"`
}

type UserResponsePagination struct {
	Status     string                `json:"status"`
	Message    string                `json:"message"`
	Results    []*User               `json:"results"`
	Pagination *presenter.Pagination `json:"pagination"`
}

type User struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"`
	Posts     []*Post `json:"post"`
	LastName  string  `json:"lastName"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type UserDetailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Results *User  `json:"results"`
}

type Post struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
