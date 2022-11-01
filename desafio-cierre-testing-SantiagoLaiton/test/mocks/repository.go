package mocks

import (
	"fmt"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/models"
)

type MockRepository struct {
	DataMock []models.Product
	Error    string
}

func (m *MockRepository) GetAll() ([]models.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}
