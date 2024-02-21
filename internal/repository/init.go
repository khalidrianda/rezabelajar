package repository

import (
	"belajar/internal/repository/content"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Content content.Repo
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		Content: content.NewContentRepo(db),
	}
}
