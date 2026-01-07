package usecase

import (
	"session-24/internal/data/entity"
	"session-24/internal/data/repository"
)

type OrderUsecase struct {
	repo *repository.OrderRepository
}

func NewOrderUsecase(repo *repository.OrderRepository) *OrderUsecase {
	return &OrderUsecase{repo: repo}
}

func (u *OrderUsecase) CreateOrder(order *entity.Order) error {
	if err := u.repo.Create(order); err != nil {
		return err
	}
	return nil
}
