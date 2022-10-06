package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var ListProducts []Product

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Color        string  `json:"color" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Stock        int     `json:"stock" validate:"required"`
	Cod          string  `json:"cod" validate:"required"`
	Published    string  `json:"published" validate:"required"`
	CreationDate string  `json:"creationDate" validate:"required"`
}
type ProductFilter struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Cod          string  `json:"cod"`
	Published    string  `json:"published"`
	CreationDate string  `json:"creationDate"`
}

func (product *Product) validate() error {
	validate := validator.New()
	return validate.Struct(product)
}

func readFile() (products []Product) {
	readFile, err := os.ReadFile("productos.json")
	if err != nil {
		panic("Hubo error en la lectura del archivo")
	}
	if err := json.Unmarshal(readFile, &products); err != nil {
		log.Fatal(err)
	}
	return products
}

func ProductFilters(c *gin.Context) {
	filter := ProductFilter{}
	id, haveId := c.GetQuery("id")
	name, haveName := c.GetQuery("name")
	color, haveColor := c.GetQuery("color")
	price, havePrice := c.GetQuery("price")
	stock, haveStock := c.GetQuery("stock")
	cod, haveCod := c.GetQuery("cod")
	published, havePublished := c.GetQuery("published")
	creationDate, haveCreationDate := c.GetQuery("creationDate")
	if haveId && id != "" {
		filter.Id, _ = strconv.Atoi(id)
	}
	if haveName {
		filter.Name = name
	}
	if haveColor {
		filter.Color = color
	}
	if havePrice && price != "" {
		filter.Price, _ = strconv.ParseFloat(price, 64)
	}
	if haveStock && stock != "" {
		filter.Stock, _ = strconv.Atoi(stock)
	}
	if haveCod {
		filter.Cod = cod
	}
	if havePublished {
		filter.Published = published
	}
	if haveCreationDate {
		filter.CreationDate = creationDate
	}
	productsFilter := FilterProducts(filter)
	if len(productsFilter) > 0 {
		c.JSON(http.StatusOK, productsFilter)
	} else {
		c.String(http.StatusNoContent, "No se encontro ningun producto")
	}
}

func FilterProducts(filter ProductFilter) (productsFilter []Product) {
	products := readFile()
	for _, product := range products {
		var equal = true
		if filter.Id != 0 && filter.Id != product.Id {
			equal = false
		}
		if filter.Name != "" && filter.Name != product.Name {
			equal = false
		}
		if filter.Color != "" && filter.Color != product.Color {
			equal = false
		}
		if filter.Price != 0.0 && filter.Price != product.Price {
			equal = false
		}
		if filter.Stock != 0 && filter.Stock != product.Stock {
			equal = false
		}
		if filter.Cod != "" && filter.Cod != product.Cod {
			equal = false
		}
		if filter.Published != "" && filter.Published != product.Published {
			equal = false
		}
		if filter.CreationDate != "" && filter.CreationDate != product.CreationDate {
			equal = false
		}
		if equal {
			productsFilter = append(productsFilter, product)
		}
		fmt.Println(product.Id, equal)
	}
	return productsFilter
}
func GetProductById(c *gin.Context) {
	products := readFile()
	id, _ := strconv.Atoi(c.Param("id"))

	var product Product
	for _, value := range products {
		if value.Id == id {
			product = value
		}
	}
	if product.Id != 0 {
		c.JSON(http.StatusOK, product)
	} else {
		c.String(http.StatusNoContent, "No se encontro ningun producto")
	}

}
func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Santiago!",
		})
	})

	router.GET("/products", func(ctx *gin.Context) {
		products := readFile()
		ctx.JSON(200, gin.H{
			"Products": products,
		})

	})

	router.GET("/productFilters", ProductFilters)

	router.GET("/product/:id", GetProductById)

	router.POST("/addProducto", AddProduc)

	router.Run()
}

func AddProduc(ctx *gin.Context) {
	var NewProduct Product
	err := ctx.ShouldBindJSON(&NewProduct)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := NewProduct.validate(); err != nil {
		validateErrors := err.(validator.ValidationErrors)
		r := []string{}
		for _, e := range validateErrors {
			r = append(r, "el campo"+e.Field()+" es requerido")
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": r})
		return
	}
	id := len(ListProducts)
	NewProduct.Id = id + 1
	ListProducts = append(ListProducts, NewProduct)
	if NewProduct.Id != 0 {
		ctx.JSON(http.StatusOK, NewProduct)
	} else {
		ctx.String(http.StatusNoContent, "no dejo registrar el producto")
	}
}
