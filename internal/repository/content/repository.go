package content

import (
	"belajar/internal/model"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	DB *sqlx.DB
}

func NewContentRepo(db *sqlx.DB) Repo {
	return Repo{
		DB: db,
	}
}

func (rq *Repo) GetData(query string, author string) ([]model.Article, error) {
	var result []model.Article

	qGetData := "SELECT id, author, COALESCE(body, ''), title FROM articles "

	rows, err := rq.DB.DB.Query(qGetData)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var article model.Article
		err := rows.Scan(&article.ID, &article.Author, &article.Body, &article.Title)
		if err != nil {
			return result, err
		}
		result = append(result, article)
	}

	return result, nil
}

func (rq *Repo) Insert(data model.Article) (int, error) {
	var result int

	qInsert := "INSERT INTO articles(author, title, body) VALUES ($1, $2, $3) returning id"

	if err := rq.DB.DB.QueryRow(qInsert, data.Author, data.Title, data.Body).Scan(&result); err != nil {
		return result, err
	}

	return result, nil
}
