package products

import (
	"strconv"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/internal/intities"
)

type Service interface {
	NewProduct(name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error)
	GetProductById(id string) (intities.Product, error)
	GetAll() (list []intities.Product, err error)
	UpdateProduct(id int, name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) NewProduct(name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error) {
	product, err := s.repository.NewProduct(name, color, price, stock, cod, published, creationDate)
	if err != nil {
		return intities.Product{}, err
	}
	return product, nil
}
func (s *service) GetProductById(id string) (intities.Product, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return intities.Product{}, err
	}
	product, err := s.repository.GetProductById(idInt)
	if err != nil {
		return intities.Product{}, err
	}
	return product, nil
}
func (s *service) GetAll() (list []intities.Product, err error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return list, err
	}
	return products, nil
}
func (s *service) UpdateProduct(id int, name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error) {
	return s.repository.UpdateProduct(id, name, color, price, stock, cod, published, creationDate)
}
func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
