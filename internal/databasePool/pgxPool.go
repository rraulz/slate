package databasePool

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	PgxPool *pgxpool.Pool
}

func NewDatabase(ctx context.Context, connString string) (pgInstance *Database, err error) {

	var db *pgxpool.Pool
	var pgOnce sync.Once

	pgOnce.Do(func() {
		db, err = pgxpool.New(ctx, connString)
		if err != nil {
			fmt.Println("Unable to connect to database")
			return
		}

		err = db.Ping(ctx)
		if err != nil {
			fmt.Println("Unable to ping database", err)
			return
		}
		fmt.Println("Database connected")

		pgInstance = &Database{db}
	})
	return
}

func (pg *Database) Ping(ctx context.Context) error {
	return pg.PgxPool.Ping(ctx)
}

func (pg *Database) Close() {
	pg.PgxPool.Close()
}

///////////////////

func (pg *Database) InsertRow(ctx context.Context, query string, args pgx.NamedArgs) error {

	_, err := pg.PgxPool.Exec(ctx, query, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (pg *Database) FetchRows(ctx context.Context, query string) (rows pgx.Rows, err error) {

	rows, err = pg.PgxPool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("unable to query users: %w", err)
	}
	return
}

func (pg *Database) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return pg.PgxPool.QueryRow(ctx, query, args...)
}
