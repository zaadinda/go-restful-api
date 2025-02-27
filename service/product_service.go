package service

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) (web.ProductResponse, error)
	Update(ctx context.Context, request web.ProductUpdateRequest) (web.ProductResponse, error)
	Delete(ctx context.Context, productId string) error
	FindById(ctx context.Context, productId string) (web.ProductResponse, error)
	FindAll(ctx context.Context) ([]web.ProductResponse, error)
}
