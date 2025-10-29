package handler

import (
	"balance-api/internal/domain/models"
	"balance-api/internal/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BalanceHandler struct {
	service *service.BalanceService
}

func NewBalanceHandler(service *service.BalanceService) *BalanceHandler {
	return &BalanceHandler{service: service}
}

func (h *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	userID := vars["user_id"]

	response, err := h.service.GetBalance(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *BalanceHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.WithdrawResponse{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	response, err := h.service.Withdraw(&req)
	if err != nil {
		statusCode := http.StatusBadRequest
		if response.Message == "User not found" {
			statusCode = http.StatusNotFound
		} else if response.Message == "Insufficient funds" {
			statusCode = http.StatusBadRequest
		}
		
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}