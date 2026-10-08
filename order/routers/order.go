package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order/models"
	"order/repository"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type RouterOrder struct {
	orderRepository *repository.OrderRepository
}

func NewRouterOrder(orderRepository *repository.OrderRepository) *RouterOrder {
	return &RouterOrder{
		orderRepository: orderRepository,
	}
}

func (ro *RouterOrder) OrderRouters() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getOrder)
	r.Post("/", ro.postOrder)
	return r
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("phone", r.Body)
	w.Write([]byte("phone"))
}

func (re *RouterOrder) postOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var request Product
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: "invalid JSON",
		})

		return
	}
	fmt.Println(request)

	data := &models.Orders{
		ID:       request.ID,
		Quantity: request.Quantity,
	}

	err = re.orderRepository.Create(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: "failed to create order",
		})

		return
	}
}
