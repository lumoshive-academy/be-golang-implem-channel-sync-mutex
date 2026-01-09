package usecase

import (
	"session-24/internal/data/entity"
	"session-24/internal/data/repository"
	"session-24/pkg/utils"
)

type OrderUsecase struct {
	repo     *repository.OrderRepository
	emailJob chan<- utils.EmailJob
}

func NewOrderUsecase(repo *repository.OrderRepository, emailJob chan<- utils.EmailJob) *OrderUsecase {
	return &OrderUsecase{repo: repo, emailJob: emailJob}
}

func (u *OrderUsecase) CreateOrder(order *entity.Order) error {
	if err := u.repo.Create(order); err != nil {
		return err
	}

	// send email
	u.emailJob <- utils.EmailJob{Email: order.Email}
	// err := utils.SendEmail(order.Email)
	// if err != nil {
	// 	return err
	// }

	return nil
}
