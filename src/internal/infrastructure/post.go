package infrastructure

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/domain"
	"github.com/sohosai/go-gin-sqlx-pg-template/internal/utils"
)

type PgPostRepository struct {
	DB *sqlx.DB
}

type PostRow struct {
	Id        string    `db:"id"`
	Text      string    `db:"text"`
	CreatedBy string    `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func intoPost(row PostRow) domain.Post {
	return domain.Post{
		Id:        row.Id,
		Text:      row.Text,
		CreatedBy: row.CreatedBy,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}

func (r PgPostRepository) GetPosts() ([]domain.Post, error) {
	postRows := []PostRow{}
	if err := r.DB.Select(&postRows, `
		SELECT 
			id, text, created_by, created_at, updated_at, deleted_at
		FROM posts`,
	); err != nil {
		log.Printf("sqlx.DB.Select error %s", err)
		return nil, err
	}

	return utils.Map(postRows, intoPost), nil
}
