package usecase

import (
	"belajar/internal/repository"
	"belajar/internal/usecase/content"
)

type Usecase struct {
	Content content.ContentUsecase
}

func NewUsecase(repo repository.Repository) Usecase {
	return Usecase{
		Content: content.NewContentUsecase(repo),
	}
}
