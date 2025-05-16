package database

import (
	"context"
	"database/sql"
)

type postgresTxKey struct{}

type Transactor interface {
	Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error
}

type DBExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) Transactor {
	return &transactor{db: db}
}

func (t *transactor) Transaction(ctx context.Context, txFunc func(txCtx context.Context) error) error {
	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	err = txFunc(injectTx(ctx, tx))
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit()
}

func injectTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, postgresTxKey{}, tx)
}

func extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(postgresTxKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}

type PostgresWrapper struct {
	db *sql.DB
}

func NewPostgresWrapper(db *sql.DB) *PostgresWrapper {
	return &PostgresWrapper{
		db: db,
	}
}

func (w *PostgresWrapper) Start(ctx context.Context) DBExecutor {
	tx := extractTx(ctx)
	if tx != nil {
		return tx
	}
	return w.db
}
