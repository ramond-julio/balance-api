package models

import "time"

type Balance struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WithdrawRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type WithdrawResponse struct {
	Success      bool    `json:"success"`
	Message      string  `json:"message"`
	NewBalance   float64 `json:"new_balance,omitempty"`
	TransactionID int64  `json:"transaction_id,omitempty"`
}

type BalanceResponse struct {
	Success bool    `json:"success"`
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
	Message string  `json:"message,omitempty"`
}

type Transaction struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Type      string    `json:"type"` // "withdraw", "deposit"
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}