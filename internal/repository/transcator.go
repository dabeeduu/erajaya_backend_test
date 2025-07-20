package repository

import (
	"backend_golang/utils"
	"context"
	"database/sql"
)

type Transactor interface {
	WithinTransaction(ctx context.Context, txFunc func(context.Context) error) error
}

type transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *transactor {
	return &transactor{db: db}
}

func (t *transactor) WithinTransaction(ctx context.Context, txFunc func(context.Context) error) error {
	tx, err := t.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})

	if err != nil {
		return err
	}

	defer func() {
		tx.Rollback()
	}()

	txCtx := utils.TxToContext(ctx, tx)
	if err := txFunc(txCtx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
