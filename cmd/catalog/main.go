package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Ghalbavieira/e-commerce.git/internal/database"
	"github.com/Ghalbavieira/e-commerce.git/internal/services"
	"github.com/Ghalbavieira/e-commerce.git/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/e-commerce")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := services.NewCategoryService(*categoryDB)
	
	productDB := database.NewProductDB(db)
	productService := services.NewProductService(*productDB)

	WebCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	WebProductHandler := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", WebCategoryHandler.GetCategories)
	c.Get("/category/{id}", WebCategoryHandler.GetCategories)
	c.Post("/category", WebCategoryHandler.CreateCategory)

	c.Get("/product/{id}", WebProductHandler.GetProduct)
	c.Get("/product/", WebProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", WebProductHandler.GetProductsByCategoryID)
	c.Post("/product", WebProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}