package mocks

import (
	"fmt"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/products"
)

type MockService struct {
	DataMock []products.Product
	Error    string
}

func (m *MockService) GetAll() ([]products.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}
