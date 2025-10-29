package repository

import "balance-api/internal/domain/models"	

type BalanceRepository interface {
	GetBalanceByUserID(userID string) (*models.Balance, error)
	UpdateBalance(userID string, newBalance float64) error
	CreateTransaction(transaction *models.Transaction) (int64, error)
	BeginTransaction() (interface{}, error)
	Commit(tx interface{}) error
	Rollback(tx interface{}) error
}