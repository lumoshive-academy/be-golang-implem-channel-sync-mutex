package dto

type CreateOrderRequest struct {
	Email string `json:"email"`
	Total int64  `json:"total"`
}
