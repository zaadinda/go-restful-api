package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type DiscountRepository interface {
	Save(ctx context.Context, discount domain.Discount) (domain.Discount, error)
	Update(ctx context.Context, discount domain.Discount) (domain.Discount, error)
	Delete(ctx context.Context, discount domain.Discount) error
	FindById(ctx context.Context, discountId string) (domain.Discount, error)
	FindAll(ctx context.Context) ([]domain.Discount, error)
}

type DiscountRepositoryImpl struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &DiscountRepositoryImpl{db: db}
}

func (repository *DiscountRepositoryImpl) Save(ctx context.Context, discount domain.Discount) (domain.Discount, error) {
	if err := repository.db.WithContext(ctx).Create(&discount).Error; err != nil {
		return domain.Discount{}, err
	}
	return discount, nil
}

func (repository *DiscountRepositoryImpl) Update(ctx context.Context, discount domain.Discount) (domain.Discount, error) {
	if err := repository.db.WithContext(ctx).Save(&discount).Error; err != nil {
		return domain.Discount{}, err
	}
	return discount, nil
}

func (repository *DiscountRepositoryImpl) Delete(ctx context.Context, discount domain.Discount) error {
	return repository.db.WithContext(ctx).Delete(&discount).Error
}

func (repository *DiscountRepositoryImpl) FindById(ctx context.Context, discountId string) (domain.Discount, error) {
	var discount domain.Discount
	err := repository.db.WithContext(ctx).First(&discount, "employee_id = ?", discountId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return discount, errors.New("employee not found")
	}
	return discount, err
}

func (repository *DiscountRepositoryImpl) FindAll(ctx context.Context) ([]domain.Discount, error) {
	var discounts []domain.Discount
	return discounts, repository.db.WithContext(ctx).Find(&discounts).Error
}
