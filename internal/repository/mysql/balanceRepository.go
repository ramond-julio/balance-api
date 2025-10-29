package mysql

import (
	"balance-api/internal/domain/models"
	"database/sql"
	"fmt"
)

type BalanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) *BalanceRepository {
	return &BalanceRepository{db: db}
}

func (r *BalanceRepository) GetBalanceByUserID(userID string) (*models.Balance, error) {
	query := `SELECT id, user_id, amount, created_at, updated_at FROM balances WHERE user_id = ?`
	
	var balance models.Balance
	err := r.db.QueryRow(query, userID).Scan(
		&balance.ID,
		&balance.UserID,
		&balance.Amount,
		&balance.CreatedAt,
		&balance.UpdatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	
	if err != nil {
		return nil, err
	}
	
	return &balance, nil
}

func (r *BalanceRepository) UpdateBalance(userID string, newBalance float64) error {
	query := `UPDATE balances SET amount = ?, updated_at = NOW() WHERE user_id = ?`
	
	result, err := r.db.Exec(query, newBalance, userID)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	
	return nil
}

func (r *BalanceRepository) CreateTransaction(transaction *models.Transaction) (int64, error) {
	query := `INSERT INTO transactions (user_id, type, amount, created_at) VALUES (?, ?, ?, NOW())`
	
	result, err := r.db.Exec(query, transaction.UserID, transaction.Type, transaction.Amount)
	if err != nil {
		return 0, err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	
	return id, nil
}

func (r *BalanceRepository) BeginTransaction() (interface{}, error) {
	return r.db.Begin()
}

func (r *BalanceRepository) Commit(tx interface{}) error {
	return tx.(*sql.Tx).Commit()
}

func (r *BalanceRepository) Rollback(tx interface{}) error {
	return tx.(*sql.Tx).Rollback()
}