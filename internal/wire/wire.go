package wire

import (
	handler "session-24/internal/adaptor"
	"session-24/internal/data/repository"
	"session-24/internal/usecase"
	"session-24/pkg/utils"
	"sync"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Route *chi.Mux
	Stop  chan struct{}
	WG    *sync.WaitGroup
}

func Wiring(repo *repository.OrderRepository) *App {
	r := chi.NewRouter()

	emailJobs := make(chan utils.EmailJob, 10) // BUFFER
	stop := make(chan struct{})
	metrics := &utils.Metrics{}
	wg := &sync.WaitGroup{}

	utils.StartEmailWorkers(3, emailJobs, stop, metrics, wg)

	wireOrder(r, repo, emailJobs)

	return &App{
		Route: r,
		Stop:  stop,
		WG:    wg,
	}
}

func wireOrder(route *chi.Mux, repo *repository.OrderRepository, emailJob chan<- utils.EmailJob) {
	uc := usecase.NewOrderUsecase(repo, emailJob)
	h := handler.NewOrderHandler(uc)
	route.Post("/orders", h.Create)
}
