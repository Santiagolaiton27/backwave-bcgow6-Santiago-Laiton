package storage

import (
	"context"
	"log"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/internal/domain"
)

type Service interface {
	Get(ctx context.Context, name string) (domain.Product, error)
	Save(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.Product, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Get(ctx context.Context, name string) (product domain.Product, err error) {
	product, err = s.repo.Get(ctx, name)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
func (s *service) Save(ctx context.Context, m domain.Product) (domain.Product, error) {
	// if s.repo.Exists(ctx, m.ID) {
	// 	return domain.Movie{}, errors.New("error: movie id already exists")
	// }
	movie_id, err := s.repo.Save(ctx, m)
	if err != nil {
		return domain.Product{}, err
	}

	m.Id = int(movie_id)
	return m, nil
}
func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repo.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}
