package databasePool

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pgxPool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connString string) (pgInstance *Database, err error) {

	var db *pgxpool.Pool
	var pgOnce sync.Once

	pgOnce.Do(func() {
		db, err = pgxpool.New(ctx, connString)
		if err != nil {
			return
		}
		pgInstance = &Database{db}
	})
	return
}

func (pg *Database) Ping(ctx context.Context) error {
	return pg.pgxPool.Ping(ctx)
}

func (pg *Database) Close() {
	pg.pgxPool.Close()
}

///////////////////

func (pg *Database) InsertRow(ctx context.Context, query string, args pgx.NamedArgs) error {

	_, err := pg.pgxPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (pg *Database) FetchRows(ctx context.Context, query string) (rows pgx.Rows, err error) {

	rows, err = pg.pgxPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	return
}
