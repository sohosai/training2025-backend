package models

import (
	"time"

	"github.com/sohosai/go-gin-sqlx-pg-template/internal/domain"
)

type PostModel struct {
	Id        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func IntoPostModel(post domain.Post) PostModel {
	return PostModel{
		Id:        post.Id,
		Text:      post.Text,
		CreatedBy: post.CreatedBy,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		DeletedAt: post.DeletedAt,
	}
}
