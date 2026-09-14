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

	err = validatePassword(request.Password)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, "successfully")
	fmt.Fprintln(w, "hashed password", hashedPassword)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MessageResponse{Message: "user registered"})

}

func validatePhone(phone string) error {
	if len(phone) != 11 {
		return fmt.Errorf("phone number must be 11 digits")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	hasDigit := false

	for _, r := range password {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}

	if hasDigit == false {
		return fmt.Errorf("There must be at least one digit.")
	}

	return nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
