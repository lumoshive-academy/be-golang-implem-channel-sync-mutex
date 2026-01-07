package wire

import (
	handler "session-24/internal/adaptor"
	"session-24/internal/data/repository"
	"session-24/internal/usecase"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Route *chi.Mux
}

func Wiring(repo *repository.OrderRepository) *App {
	r := chi.NewRouter()

	wireOrder(r, repo)

	return &App{
		Route: r,
	}
}

func wireOrder(route *chi.Mux, repo *repository.OrderRepository) {
	uc := usecase.NewOrderUsecase(repo)
	h := handler.NewOrderHandler(uc)
	route.Post("/orders", h.Create)
}
