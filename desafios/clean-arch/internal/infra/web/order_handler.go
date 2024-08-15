package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"CleanArch/internal/entity"
	"CleanArch/internal/usecase"
	"CleanArch/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
	}
}

type WebListOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func NewWebListOrderHandler(
	OrderRepository entity.OrderRepositoryInterface,
) *WebListOrderHandler {
	return &WebListOrderHandler{
		OrderRepository: OrderRepository,
	}
}

func (h *WebListOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	pageInt := 0
	pageSizeInt := 100
	var err error = nil
	if len(r.URL.Query().Get("page")) > 0 && len(r.URL.Query().Get("pageSize")) > 0 {
		pageInt, err = strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			pageInt = 0
		}
		pageSizeInt, err = strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil {
			pageInt = 0
			pageSizeInt = 100
		}
	}

	dto := usecase.ListOrderInputDTO{
		Page:     pageInt,
		PageSize: pageSizeInt,
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	listOrder := usecase.NewListOrderUseCase(h.OrderRepository)
	output, err := listOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
