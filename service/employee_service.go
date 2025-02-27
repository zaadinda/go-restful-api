package service

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

type EmployeeService interface {
	Create(ctx context.Context, request web.EmployeeCreateRequest) (web.EmployeeResponse, error)
	Update(ctx context.Context, request web.EmployeeUpdateRequest) (web.EmployeeResponse, error)
	Delete(ctx context.Context, employeeId string) error
	FindById(ctx context.Context, employeeId string) (web.EmployeeResponse, error)
	FindAll(ctx context.Context) ([]web.EmployeeResponse, error)
}
