package handler

import (
	"belajar/internal/handler/content"
	"belajar/internal/usecase"
)

type Handler struct {
	Content content.ContentHandler
}

func NewHandler(usecase usecase.Usecase) Handler {
	return Handler{
		Content: content.NewContentHandler(usecase),
	}
}
