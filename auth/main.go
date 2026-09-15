package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func main() {
	r := chi.NewRouter()

	r.Use(loggerMiddleware)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/auth/reg", registerHandler)
	})

	fmt.Println("server started on : 8000")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		fmt.Println(err)
	}
}

//	func rootHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintln(w, map[string]string{"message": "ok"})
//	}
func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	err = validatePhone(request.Phone)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

