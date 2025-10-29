package service

import (
	"balance-api/internal/domain/models"
	"balance-api/internal/repository"
	"fmt"
)

type BalanceService struct {
	repo repository.BalanceRepository
}

func NewBalanceService(repo repository.BalanceRepository) *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetBalance(userID string) (*models.BalanceResponse, error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	balance, err := s.repo.GetBalanceByUserID(userID)
	if err != nil {
		return &models.BalanceResponse{
			Success: false,
			UserID:  userID,
			Message: err.Error(),
		}, err
	}

	return &models.BalanceResponse{
		Success: true,
		UserID:  balance.UserID,
		Balance: balance.Amount,
	}, nil
}

func (s *BalanceService) Withdraw(req *models.WithdrawRequest) (*models.WithdrawResponse, error) {
	// Validation
	if req.UserID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	if req.Amount <= 0 {
		return nil, fmt.Errorf("withdrawal amount must be greater than zero")
	}

	// Get current balance
	balance, err := s.repo.GetBalanceByUserID(req.UserID)
	if err != nil {
		return &models.WithdrawResponse{
			Success: false,
			Message: "User not found",
		}, err
	}

	// Check sufficient funds
	if balance.Amount < req.Amount {
		return &models.WithdrawResponse{
			Success: false,
			Message: "Insufficient funds",
		}, fmt.Errorf("insufficient funds")
	}

	// Calculate new balance
	newBalance := balance.Amount - req.Amount

	// Update balance
	err = s.repo.UpdateBalance(req.UserID, newBalance)
	if err != nil {
		return &models.WithdrawResponse{
			Success: false,
			Message: "Failed to update balance",
		}, err
	}

	// Create transaction record
	transaction := &models.Transaction{
		UserID: req.UserID,
		Type:   "withdraw",
		Amount: req.Amount,
	}

	transactionID, err := s.repo.CreateTransaction(transaction)
	if err != nil {
		// Log error but don't fail the withdrawal
		fmt.Printf("Failed to create transaction record: %v\n", err)
	}

	return &models.WithdrawResponse{
		Success:       true,
		Message:       "Withdrawal successful",
		NewBalance:    newBalance,
		TransactionID: transactionID,
	}, nil
}