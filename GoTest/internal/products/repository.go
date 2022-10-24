package products

import (
	"errors"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/GoTest/internal/intities"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/GoTest/internal/pkg/store"
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

//var ps Archivo

type Repository interface {
	NewProduct(name string, color string, price float64, stock int, cod string, published bool, creationDate string) (intities.Product, error)
	GetProductById(id int) (intities.Product, error)
	GetAll() (list []*intities.Product, err error)
	UpdateProduct(id int, name string, color string, price float64, stock int, cod string, published bool, creationDate string) (intities.Product, error)
	UpdateProductName(id int, name string) (*intities.Product, error)
	Delete(id int) error
}
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() (list []*intities.Product, err error) {
	var products []*intities.Product
	err = r.db.Read(&products)
	if err != nil {
		return []*intities.Product{}, err
	}
	return products, nil
}

func (r *repository) UpdateProduct(id int, name string, color string, price float64, stock int, cod string, published bool, creationDate string) (intities.Product, error) {
	var newProduct intities.Product
	var update bool
	var products []*intities.Product
	err := r.db.Read(&products)
	if err != nil {
		return intities.Product{}, err
	}
	for i := range products {
		if products[i].Id == id {
			newProduct = intities.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Cod: cod, Published: published, CreationDate: creationDate}
			products[i] = &newProduct
			update = true
		}
	}
	if !update {
		return intities.Product{}, errors.New("no se pudo actualizar el producto")
	}
	err = r.db.Write(&products)
	if err != nil {
		return intities.Product{}, err
	}
	return newProduct, nil
}
func (r *repository) Delete(id int) error {
	var products []*intities.Product
	err := r.db.Read(&products)
	delete := false
	pos := 0
	if err != nil {
		return err
	}
	for i := range products {
		if products[i].Id == id {
			delete = true
			pos = i
		}
	}
	if !delete {
		return errors.New("no se encontro el elemento a eliminar")
	} else {
		products = append(products[:pos], products[pos+1:]...)
	}
	err = r.db.Write(&products)
	if err != nil {
		return errors.New("error al escribir el nuevo archivo")
	}
	return nil
}
func (r *repository) NewProduct(name string, color string, price float64, stock int, cod string, published bool, creationDate string) (intities.Product, error) {
	var products []intities.Product
	err := r.db.Read(&products)
	if err != nil {
		return intities.Product{}, err
	}
	id := +1
	product := intities.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Cod: cod, Published: published, CreationDate: creationDate}
	products = append(products, product)
	err = r.db.Write(&products)
	if err != nil {
		return intities.Product{}, err
	}
	return product, nil
}
func (r *repository) GetProductById(id int) (intities.Product, error) {
	var product intities.Product
	var products []intities.Product
	err := r.db.Read(&products)
	if err != nil {
		return intities.Product{}, err
	}
	for _, value := range products {
		if value.Id == id {
			product = value
		}
	}
	if (intities.Product{} == product) {
		return product, errors.New("no se encontro ningun producto con ese id")
	}
	return product, nil
}

func (r *repository) UpdateProductName(id int, name string) (*intities.Product, error) {

	var update bool
	var products []*intities.Product
	err := r.db.Read(&products)
	if err != nil {
		return nil, err
	}
	var newProduct *intities.Product
	for _, value := range products {
		if value.Id == id {
			value.Name = name
			newProduct = value
			update = true
		}
	}
	if !update {
		return nil, errors.New("no se pudo actualizar el producto")
	}
	err = r.db.Write(&products)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}
