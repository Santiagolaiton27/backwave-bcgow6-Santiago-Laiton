package mocks

import (
	"fmt"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/models"
)

type MockStore struct {
	DataMock []models.Product
	ErrWrite string
	ErrRead  string
}

func (m *MockStore) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]models.Product)
	*a = m.DataMock
	return nil
}
func (m *MockStore) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]models.Product)
	m.DataMock = append(m.DataMock, a[len(a)-1])
	return nil
}
