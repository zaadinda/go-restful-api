package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := repository.db.WithContext(ctx).Create(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := repository.db.WithContext(ctx).Save(&product).Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product domain.Product) error {
	return repository.db.WithContext(ctx).Delete(&product).Error
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId string) (domain.Product, error) {
	var product domain.Product
	err := repository.db.WithContext(ctx).First(&product, "product_id = ?", productId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return product, errors.New("product not found")
	}
	return product, err
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	return products, repository.db.WithContext(ctx).Find(&products).Error
}
