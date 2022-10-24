package products

import (
	"fmt"
	"testing"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/GoTest/internal/intities"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	Products []intities.Product
}

func (stubStore *StubStore) Write(data interface{}) error {
	return nil
}
func (stubStore *StubStore) Read(data interface{}) error {
	a := data.(*[]*intities.Product)
	p1 := &intities.Product{
		Id:           8,
		Name:         "Televisor",
		Color:        "Negro",
		Price:        2000000,
		Stock:        20,
		Cod:          "Tele01",
		Published:    true,
		CreationDate: "05-10-2022",
	}
	p2 := &intities.Product{
		Id:           9,
		Name:         "Monitor",
		Color:        "Negro",
		Price:        2000000,
		Stock:        20,
		Cod:          "Moni01",
		Published:    true,
		CreationDate: "05-10-2022",
	}

	*a = append(*a, p1)
	*a = append(*a, p2)

	return nil
}

func TestGetAll(t *testing.T) {
	db := StubStore{}
	repository := NewRepository(&db)
	expected := []*intities.Product{
		{
			Id:           8,
			Name:         "Televisor",
			Color:        "Negro",
			Price:        2000000,
			Stock:        20,
			Cod:          "Tele01",
			Published:    true,
			CreationDate: "05-10-2022",
		},
		{
			Id:           9,
			Name:         "Monitor",
			Color:        "Negro",
			Price:        2000000,
			Stock:        20,
			Cod:          "Moni01",
			Published:    true,
			CreationDate: "05-10-2022",
		},
	}
	result, err := repository.GetAll()
	assert.Nil(t, err)
	fmt.Println(result)
	assert.Equal(t, expected, result)
}

type MockStore struct {
	ReadInvoked bool
	Data        []*intities.Product
}

func (fs *MockStore) Read(data interface{}) error {
	fs.ReadInvoked = true
	a := data.(*[]*intities.Product)
	*a = fs.Data
	return nil
}

func (fs *MockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {
	id, name := 9, "Santiago"
	products := []*intities.Product{{
		Id:           9,
		Name:         "Monitor",
		Color:        "Negro",
		Price:        2000000,
		Stock:        20,
		Cod:          "Moni01",
		Published:    true,
		CreationDate: "05-10-2022",
	}}
	mock := MockStore{Data: products}
	repo := NewRepository(&mock)
	productUploaded, err := repo.UpdateProductName(id, name)
	assert.Nil(t, err)
	assert.Equal(t, id, productUploaded.Id)
	assert.Equal(t, name, productUploaded.Name)

}
