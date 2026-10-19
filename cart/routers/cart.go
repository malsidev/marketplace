package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"order/models"
	"order/repository"
	"order/service"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	UserID    int `json:"UserID"`
	ProductID int `json:"ProductID"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	UserID int `json:"UserID"`
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
	r.Get("/", ro.getCart)
	r.Post("/", ro.postCart)
	return r
}

func (re *RouterCart) getCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request Cart
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, `{"message":"invalid JSON"}`, http.StatusBadRequest)
		return
	}

	items, err := re.cartRepository.GetCart(request.UserID)
	if err != nil {
		http.Error(w, `{"message":"failed to get cart"}`, http.StatusInternalServerError)
		return
	}

	productItems := make([]service.CartItemRequest, 0, len(items))

	for _, item := range items {
		productItems = append(productItems, service.CartItemRequest{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	product, err := service.GetProduct(productItems)
	if err != nil {
		http.Error(w, `{"message":"seeervice"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
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
		ProductID: request.ProductID,
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
