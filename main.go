package main

import (
	"belajar/config.go"
	"belajar/internal/handler"
	"belajar/internal/repository"
	"belajar/internal/usecase"
	"belajar/utils/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.NewConfig()               // 1
	initDB := db.NewDBConnection(conf)       //2
	repo := repository.NewRepository(initDB) //3 //1
	usecase := usecase.NewUsecase(repo)      // 4  //2
	handler := handler.NewHandler(usecase)   // 5  //3

	e := fiber.New()
	api := e.Group("/api")

	api.Get("/getdata", handler.Content.GetData)
	api.Post("/insert", handler.Content.Insert)

	e.Listen(":8000")
}
