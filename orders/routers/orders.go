package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func orderRouters() http.Handler {
	r := chi.NewRouter()

	r.Get("/", orderPost)

	return r
}

func orderPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
