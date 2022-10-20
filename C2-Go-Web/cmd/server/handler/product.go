package handler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/internal/products"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	Name         string  `json:"name" validate:"required"`
	Color        string  `json:"color" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
	Stock        int     `json:"stock" validate:"required"`
	Cod          string  `json:"cod" validate:"required"`
	Published    *bool   `json:"published" validate:"required"`
	CreationDate string  `json:"creationDate" validate:"required"`
}
type ProductFilters struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Cod          string  `json:"cod"`
	Published    string  `json:"published"`
	CreationDate string  `json:"creationDate"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

func (p *Product) GetProductById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAccepted := validateToken(ctx.GetHeader("Token"))
		if !tokenAccepted {
			ctx.JSON(http.StatusUnauthorized, "token invalido")
			return
		}
		if id := ctx.Param("id"); id != "" {
			p, err := p.service.GetProductById(id)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(http.StatusOK, web.NewResponse(200, p, ""))
			}
		} else {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "No tiene el id para buscar el producto"))
		}
	}
}

func (p *Product) SaveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAccepted := validateToken(ctx.GetHeader("Token"))
		if !tokenAccepted {
			ctx.JSON(http.StatusUnauthorized, "token invalido")
			return
		}
		var request Request
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// if err := request.validate(); err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{
		// 		"error": err.Error(),
		// 	})
		// 	return
		// }
		if request.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el nombre"))
			return
		}
		if request.Color == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el color"))
			return
		}
		if request.Price == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar un precio diferente de 0"))
			return
		}
		if request.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar las cantidades diponibles y tienen que ser diferente de 0"))
			return
		}
		if request.Cod == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el Codigo"))
			return
		}
		if request.Published == nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el Codigo"))
			return
		}
		if request.CreationDate == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar la fecha de creacion"))
			return
		}
		p, err := p.service.NewProduct(request.Name, request.Color, request.Price, request.Stock, request.Cod, *request.Published, request.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}
func (request *Request) validate() error {
	validate := validator.New()
	return validate.Struct(request)
}
func (c *Product) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAccepted := validateToken(ctx.GetHeader("Token"))
		if !tokenAccepted {
			ctx.JSON(http.StatusUnauthorized, "token invalido")
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "El id debe ser un numero")
			return
		}
		var request Request
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		// if err := request.validate(); err != nil {
		// 	ctx.JSON(http.StatusOK, err.Error())
		// 	return
		// }
		if request.Name == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el nombre"))
			return
		}
		if request.Color == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el color"))
			return
		}
		if request.Price == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar un precio diferente de 0"))
			return
		}
		if request.Stock == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar las cantidades diponibles y tienen que ser diferente de 0"))
			return
		}
		if request.Cod == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar el Codigo"))
			return
		}

		if request.CreationDate == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "Para añadir el producto tiene que ingresar la fecha de creacion"))
			return
		}
		p, err := c.service.UpdateProduct(id, request.Name, request.Color, request.Price, request.Stock, request.Cod, *request.Published, request.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}
func validateToken(token string) (ok bool) {
	correct := true
	if token != os.Getenv("TOKEN") {
		correct = false
	}
	return correct
}
func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAccepted := validateToken(ctx.GetHeader("Token"))
		if !tokenAccepted {
			ctx.JSON(http.StatusUnauthorized, "token invalido")
			return
		}
		products, err := p.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		if len(products) == 0 {
			ctx.JSON(http.StatusNoContent, "No hay ningun producto creado")
			return
		} else {
			ctx.JSON(http.StatusOK, products)
			return
		}
	}
}
func (p *Product) ProductFilters() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// filter := ProductFilters{}
		// id, haveId := ctx.GetQuery("id")
		// nama, haveName := ctx.GetQuery("name")

		// filter.Id = strconv.Itoa(id)

		// p.service.ProductFilters(filter)
	}
}

// func (p *Product) UpdateFields() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		tokenAccepted := validateToken(ctx.GetHeader("token"))
// 		if !tokenAccepted {
// 			ctx.JSON(http.StatusUnauthorized, "token invalido")
// 			return
// 		}
// 		id, err := strconv.Atoi(ctx.Param("id"))
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, "El id tiene que ser tipo entero")
// 			return
// 		}
// 		var change Request
// 		if err := ctx.Bind(&change); err != nil {
// 			ctx.JSON(http.StatusBadRequest, err.Error())
// 			return
// 		}
// 		p, err := p.service.UpdateFields(id, c)
// 		if err != nil {
// 			ctx.JSON(http.StatusBadRequest, err.Error())
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, p)
// 	}
// }

func (p *Product) DeleteProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenAccepted := validateToken(ctx.GetHeader("Token"))
		if !tokenAccepted {
			ctx.JSON(http.StatusUnauthorized, "token invalido")
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "El id debe ser un numero")
			return
		}
		err = p.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "Eliminacion con exito")
	}
}
