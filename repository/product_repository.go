package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, category domain.Category) (domain.Category, error)
	FindById(ctx context.Context, id int) (domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
}
