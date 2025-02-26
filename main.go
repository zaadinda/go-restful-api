package main

import (
	"github.com/aronipurwanto/go-restful-api/app"
	"github.com/aronipurwanto/go-restful-api/controller"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/aronipurwanto/go-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {

	server := fiber.New()
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	app.NewRouter(server, categoryController)

	err := server.Listen(":8080")
	helper.PanicIfError(err)
}
