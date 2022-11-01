package products

import "github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/models"

type Repository interface {
	GetAllBySeller(sellerID string) ([]models.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]models.Product, error) {
	var prodList []models.Product
	prodList = append(prodList, models.Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}
