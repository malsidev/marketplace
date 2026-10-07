package routers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func OrderRouters() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getOrder)

	return r
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("phone")
	w.Write([]byte("phone"))
}
