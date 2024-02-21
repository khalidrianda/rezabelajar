package content

import (
	"belajar/internal/model"
	"belajar/internal/repository"
	"context"
)

type ContentUsecase struct {
	Repo repository.Repository
}

func NewContentUsecase(repo repository.Repository) ContentUsecase {
	return ContentUsecase{
		Repo: repo,
	}
}

func (cu *ContentUsecase) GetData(ctx context.Context, query string, author string) ([]model.Article, error) {
	var getData []model.Article

	getData, err := cu.Repo.Content.GetData(query, author)
	if err != nil {
		return getData, err
	}

	return getData, nil
}

func (cu *ContentUsecase) Insert(ctx context.Context, data model.Article) error {

	id, err := cu.Repo.Content.Insert(data)
	if err != nil || id == 0 {
		return err
	}

	return nil
}
