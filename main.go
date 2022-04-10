package main

import (
	"yahfiilham/go-rest-api/app"
	"yahfiilham/go-rest-api/controller"
	"yahfiilham/go-rest-api/repository"
	"yahfiilham/go-rest-api/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService) 

	router := httprouter.New()

	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}