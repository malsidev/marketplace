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
	CartID   int `json:"CartID"`
	UserID   int `json:"UserID"`
	Quantity int `json:"quantity"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type RouterCart struct {
	cartRepository *repository.CartRepository
}

func NewRouterCart(cartRepository *repository.CartRepository) *RouterCart {
	return &RouterCart{
		cartRepository: cartRepository,
	}
}

func (ro *RouterCart) CartRouters() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getCart)
	r.Post("/", ro.postCart)
	return r
}

func getCart(w http.ResponseWriter, r *http.Request) {

}

func (re *RouterCart) postCart(w http.ResponseWriter, r *http.Request) {
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

	cart := &models.Carts{
		UserID: request.UserID,
	}

	cart_items := &models.CartItem{
		CartID:    cart.ID,
		ProductID: request.CartID,
		Quantity:  request.Quantity,
	}

	err = re.cartRepository.CreateCartWithItem(cart, cart_items)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: "failed to create order",
		})

		return
	}
	w.WriteHeader(http.StatusCreated)
}
