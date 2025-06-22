package models

import (
	"time"

	"github.com/sohosai/go-gin-sqlx-pg-template/internal/domain"
)

type PostModel struct {
	Id        string     `json:"id"`
	Text      string     `json:"text"`
	CreatedBy string     `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreatePostRequest struct {
	Text      string `json:"text" binding:"required"`
	CreatedBy string `json:"created_by" binding:"required"`
}

func (post CreatePostRequest) Value() domain.CreatedPost {
	return domain.CreatedPost{
		Text:      post.Text,
		CreatedBy: post.CreatedBy,
	}
}

type CreatePostResponse struct {
	Id string `json:"id"`
}

func IntoPostModel(post domain.Post) PostModel {
	return PostModel{
		Id:        post.Id,
		Text:      post.Text,
		CreatedBy: post.CreatedBy,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

/* domain.Post 型の引数をPostModel( = 実際にAPIが返すJson)に変える関数の実装
→ idから取ってきたやつをここに入れられればok
func (引数) 関数名() (戻り値1,戻り値2){
処理
}
*/
