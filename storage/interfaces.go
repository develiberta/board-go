package storage

import "awesomeProject/model"

type PostStorage interface {
	GetPost(string) model.Post
	ListPost() []model.Post
	CreatePost(model.Post)
	UpdatePost(model.Post)
	DeletePost(string)
}
