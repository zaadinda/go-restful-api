package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db: db}
}

func (repository *CustomerRepositoryImpl) Save(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	if err := repository.db.WithContext(ctx).Create(&customer).Error; err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, customer domain.Customer) (domain.Customer, error) {
	if err := repository.db.WithContext(ctx).Save(&customer).Error; err != nil {
		return domain.Customer{}, err
	}
	return customer, nil
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, customer domain.Customer) error {
	return repository.db.WithContext(ctx).Delete(&customer).Error
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, customerId string) (domain.Customer, error) {
	var customer domain.Customer
	err := repository.db.WithContext(ctx).First(&customer, "customer_id = ?", customerId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return customer, errors.New("customer not found")
	}
	return customer, err
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context) ([]domain.Customer, error) {
	var customers []domain.Customer
	return customers, repository.db.WithContext(ctx).Find(&customers).Error
}
