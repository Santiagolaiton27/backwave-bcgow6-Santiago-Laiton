package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/internal/domain"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/internal/storage"
	"github.com/gin-gonic/gin"
)

type Storage struct {
	service storage.Service
}

func NewStorage(service storage.Service) *Storage {
	return &Storage{
		service: service,
	}
}

func (m *Storage) Get() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		movie, err := m.service.Get(ctx, name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, movie)
	}
}

func (m *Storage) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		product, err = m.service.Save(ctx, product)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Product": product.Name + " added"})
	}
}

func (m *Storage) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		err = m.service.Delete(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusNoContent, "")
	}
}

func (m *Storage) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := m.service.GetAll(ctx)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(http.StatusOK, products)
	}
}
