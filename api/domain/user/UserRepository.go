package user

import (
	"context"
	"slate/api/database"

	"github.com/jackc/pgx/v5"
)

func CreateUser(ctx context.Context, pg *database.DB, user User) error {

	query := `INSERT INTO "user" (id, username, password, email) VALUES (@id, @username, @password, @email)`
	args := pgx.NamedArgs{
		"id":       user.Id,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
	}

	err := pg.InsertRow(ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(ctx context.Context, pg *database.DB) ([]User, error) {

	query := `SELECT * FROM public."user"`
	rows, err := pg.FetchRows(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[User])
}
