package transactionManager

import (
	"context"
	"database/sql"
)

// TransactionManager handles database transactions
type TransactionManager struct {
	db *sql.DB
}

// RunTransaction runs a transaction and executes the provided action
func (tm *TransactionManager) RunTransaction(ctx context.Context, action func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := tm.db.Begin()
	if err != nil {
		return err
	}

	err = action(ctx, tx)
	if err != nil {
		err := tx.Rollback()
		return err
	}

	return tx.Commit()
}
