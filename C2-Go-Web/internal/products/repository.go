package products

import (
	"errors"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/internal/intities"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/C2-Go-Web/internal/pkg/store"
)

//	type Product struct {
//		Id           int     `json:"id"`
//		Name         string  `json:"name"`
//		Color        string  `json:"color"`
//		Price        float64 `json:"price"`
//		Stock        int     `json:"stock"`
//		Cod          string  `json:"cod"`
//		Published    string  `json:"published"`
//		CreationDate string  `json:"creationDate"`
//	}
type Archivo struct {
	LastId   int
	Products []intities.Product
}

var ps Archivo

type Repository interface {
	NewProduct(name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error)
	GetProductById(id int) (intities.Product, error)
	GetAll() (list []intities.Product, err error)
	UpdateProduct(id int, name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error)
	Delete(id int) error
}
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) NewProduct(name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error) {
	err := r.db.Read(&ps)
	if err != nil {
		return intities.Product{}, err
	}
	id := ps.LastId + 1
	product := intities.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Cod: cod, Published: published, CreationDate: creationDate}
	ps.Products = append(ps.Products, product)
	ps.LastId = id
	err = r.db.Write(&ps)
	if err != nil {
		return intities.Product{}, err
	}
	return product, nil
}

func (r *repository) GetProductById(id int) (intities.Product, error) {
	var product intities.Product
	err := r.db.Read(&ps)
	if err != nil {
		return intities.Product{}, err
	}
	for _, value := range ps.Products {
		if value.Id == id {
			product = value
		}
	}
	if (intities.Product{} == product) {
		return product, errors.New("No se encontro ningun producto con ese id")
	}
	return product, nil
}
func (r *repository) GetAll() (list []intities.Product, err error) {
	err = r.db.Read(&ps)
	if err != nil {
		return []intities.Product{}, err
	}
	return ps.Products, nil
}
func (r *repository) UpdateProduct(id int, name string, color string, price float64, stock int, cod string, published string, creationDate string) (intities.Product, error) {
	var newProduct intities.Product
	var update bool
	err := r.db.Read(&ps)
	if err != nil {
		return intities.Product{}, err
	}
	for i := range ps.Products {
		if ps.Products[i].Id == id {
			newProduct = intities.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Cod: cod, Published: published, CreationDate: creationDate}
			ps.Products[i] = newProduct
			update = true
		}
	}
	if !update {
		return intities.Product{}, errors.New("No se pudo actualizar el producto")
	}
	err = r.db.Write(&ps)
	if err != nil {
		return intities.Product{}, err
	}
	return newProduct, nil
}
func (r *repository) Delete(id int) error {
	err := r.db.Read(&ps)
	delete := false
	pos := 0
	if err != nil {
		return err
	}
	for i := range ps.Products {
		if ps.Products[i].Id == id {
			delete = true
			pos = i
		}
	}
	if !delete {
		return errors.New("No se encontro el elemento a eliminar")
	} else {
		ps.Products = append(ps.Products[:pos], ps.Products[pos+1:]...)
	}
	err = r.db.Write(&ps)
	if err != nil {
		return errors.New("Error al escribir el nuevo archivo")
	}
	return nil
}
