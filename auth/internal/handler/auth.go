package handler

import (
	"encoding/json"
	"net/http"

	"auth/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

type RegisterRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: "invalid JSON",
		})

		return
	}

	err = h.authService.Register(request.Phone, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(MessageResponse{
		Message: "user registered",
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request LoginRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: "invalid JSON",
		})

		return
	}

	token, err := h.authService.Login(request.Phone, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)

		json.NewEncoder(w).Encode(MessageResponse{
			Message: err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"access_token": token,
	})
}