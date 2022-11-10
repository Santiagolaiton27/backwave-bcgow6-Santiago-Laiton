package storage

import (
	"context"
	"database/sql"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/internal/domain"
)

type Repository interface {
	Get(ctx context.Context, name string) (domain.Product, error)
	Save(ctx context.Context, p domain.Product) (int64, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.Product, error)
	//Update(ctx context.Context, p domain.Product, id int) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	SAVE_PRODUCT = "INSERT INTO products (name, type,count, price) VALUES (?, ?, ?, ?);"

	GET_ALL_PRODUCTS = "SELECT id,name,type,count,price FROM products;"

	GET_PRODUCT = "SELECT id,name,type,count,price FROM products WHERE name=?;"

	UPDATE_MOVIE = "UPDATE movies SET title=?, rating=?, awards=?, length=?, genre_id=? WHERE id=?;"

	DELETE_MOVIE = "Delete from products where id =?;"

	EXIST_MOVIE = "SELECT m.id FROM products m WHERE m.id=?;"
)

func (r *repository) Get(ctx context.Context, name string) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT, name)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Save(ctx context.Context, m domain.Product) (int64, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT) //preparamos la consulta
	if err != nil {
		return 0, err
	}
	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(m.Name, m.Type, m.Count, m.Price)
	if err != nil {
		return 0, err
	}
	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
func (r *repository) Delete(ctx context.Context, id int) error {
	stm, err := r.db.Prepare(DELETE_MOVIE)
	if err != nil {
		return err
	}
	defer stm.Close()
	_, err = stm.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(GET_ALL_PRODUCTS)
	products := []domain.Product{}
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}
