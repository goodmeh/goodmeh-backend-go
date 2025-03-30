package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RollbackFunc func(ctx context.Context) error
type CommitFunc func(ctx context.Context) error

func (q *Queries) Begin(ctx context.Context) (
	*Queries,
	RollbackFunc,
	CommitFunc,
	error,
) {
	db := q.db.(*pgxpool.Pool)
	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, nil, nil, err
	}
	return q.WithTx(tx), tx.Rollback, tx.Commit, nil
}
