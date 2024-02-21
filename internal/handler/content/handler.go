package content

import (
	"belajar/internal/model"
	"belajar/internal/usecase"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ContentHandler struct {
	Usecase usecase.Usecase
}

func NewContentHandler(usecase usecase.Usecase) ContentHandler {
	return ContentHandler{
		Usecase: usecase,
	}
}

func (ch *ContentHandler) GetData(c *fiber.Ctx) error {

	query := c.Query("query")
	author := c.Query("author")

	getData, err := ch.Usecase.Content.GetData(context.Background(), query, author)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.StatusError{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(getData)
}

func (ch *ContentHandler) Insert(c *fiber.Ctx) error {

	var data model.Article
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.StatusError{Message: err.Error()})
	}

	err = ch.Usecase.Content.Insert(context.Background(), data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.StatusError{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(model.StatusError{Message: "SUKSES"})
}
