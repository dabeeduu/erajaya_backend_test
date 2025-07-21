package utils

import (
	"backend_golang/database"
	"context"
	"database/sql"
)

type txCtxKey struct{}

func TxToContext(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txCtxKey{}, tx)
}

func contextToTx(ctx context.Context) *sql.Tx {
	tx, ok := ctx.Value(txCtxKey{}).(*sql.Tx)
	if ok {
		return tx
	}
	return nil
}

func ChooseDB(ctx context.Context, db *sql.DB) database.DB {
	if tx := contextToTx(ctx); tx != nil {
		return tx
	}

	return db
}
