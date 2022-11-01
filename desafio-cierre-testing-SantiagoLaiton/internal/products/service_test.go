package products

import (
	"testing"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	var list []models.Product
	expec := models.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	}
	list = append(list, expec)
	repo := NewRepository()
	service := NewService(repo)
	p, err := service.GetAllBySeller("mock")
	assert.Nil(t, err)
	assert.Equal(t, list, p)
}
