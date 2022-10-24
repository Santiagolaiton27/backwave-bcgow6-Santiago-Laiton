package products

import (
	"testing"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/products"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	var list []products.Product
	expec := products.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}
	list = append(list, expec)
	repo := products.NewRepository()
	service := products.NewService(repo)
	p, err := service.GetAllBySeller("mock")
	assert.Nil(t, err)
	assert.Equal(t, list, p)
}
