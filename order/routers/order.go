package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func OrderRouters() chi.Router {
	r := chi.NewRouter()
	r.Get("/", getOrder)
	r.Post("/", postOrder)
	return r
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("phone", r.Body)
	w.Write([]byte("phone"))
}

func postOrder(w http.ResponseWriter, r *http.Request) {
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
}
