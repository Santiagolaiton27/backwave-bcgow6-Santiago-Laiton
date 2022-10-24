package mocks

import (
	"fmt"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/products"
)

type MockStore struct {
	DataMock []products.Product
	ErrWrite string
	ErrRead  string
}

func (m *MockStore) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]products.Product)
	*a = m.DataMock
	return nil
}
func (m *MockStore) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]products.Product)
	m.DataMock = append(m.DataMock, a[len(a)-1])
	return nil
}
