package domain

import (
	"time"
)

type Repositories struct {
	PostRepo PostRepository
}

type Post struct {
	Id        string
	Text      string
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CreatePostArgs struct {
	Text      string
	CreatedBy string
}

type PostRepository interface {
	GetPosts() ([]Post, error)
	CreatePost(args CreatePostArgs) (Post, error)
}