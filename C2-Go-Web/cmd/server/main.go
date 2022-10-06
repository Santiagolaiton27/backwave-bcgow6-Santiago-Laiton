package main

import (
	"log"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/cmd/server/handler"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/internal/pkg/store"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al traer las variables de entorno")
	}
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/product")

	pr.POST("/add", p.SaveProduct())
	pr.GET("/:id", p.GetProductById())
	pr.GET("/products", p.GetAll())
	pr.PUT("/:id", p.UpdateProduct())
	pr.GET("/productFilter", p.ProductFilters())
	//pr.PATCH("/:id", p.UpdateFields())
	pr.DELETE("/:id", p.DeleteProduct())
	r.Run()
}
