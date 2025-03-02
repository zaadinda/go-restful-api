package web

type CustomerCreateRequest struct {
	Name       string `json:"name" validate:"required,max=100,min=1"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,max=20"`
	Address    string `json:"address" validate:"required,max=255"`
	LoyaltyPts int    `json:"loyalty_pts" validate:"gte=0"`
}

type CustomerUpdateRequest struct {
	CustomerID uint64 `json:"id" validate:"required,gte=0"`
	Name       string `json:"name" validate:"required,max=100,min=1"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,max=20"`
	Address    string `json:"address" validate:"required,max=255"`
	LoyaltyPts int    `json:"loyalty_pts" validate:"gte=0"`
}

type CustomerResponse struct {
	CustomerID uint64 `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	LoyaltyPts int    `json:"loyalty_pts"`
}
