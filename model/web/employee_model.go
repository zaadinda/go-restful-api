package web

type EmployeeCreateRequest struct {
	Name      string `json:"name" validate:"required,max=100,min=1"`
	Role      string `json:"role" validate:"required,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,max=20"`
	DateHired string `json:"date_hired" validate:"required"`
}

type EmployeeUpdateRequest struct {
	EmployeeID uint64 `json:"id" validate:"required,gte=0"`
	Name       string `json:"name" validate:"required,max=100,min=1"`
	Role       string `json:"role" validate:"required,max=50"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,max=20"`
	DateHired  string `json:"date_hired" validate:"required"`
}

type EmployeeResponse struct {
	EmployeeID uint64 `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateHired  string `json:"date_hired"`
}
