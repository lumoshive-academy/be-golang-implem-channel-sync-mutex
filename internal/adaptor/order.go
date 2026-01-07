package handler

import (
	"encoding/json"
	"net/http"
	"session-24/internal/data/entity"
	"session-24/internal/dto"
	"session-24/internal/usecase"
)

type OrderHandler struct {
	uc *usecase.OrderUsecase
}

func NewOrderHandler(uc *usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{uc: uc}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateOrderRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	order := &entity.Order{
		Email: req.Email,
		Total: req.Total,
	}

	if err := h.uc.CreateOrder(order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("order created"))
}
