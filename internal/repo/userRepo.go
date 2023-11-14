package repo

import (
	"context"
	"fmt"
	"slate/internal/databasePool"
	"slate/internal/domain"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	database *databasePool.Database
}

func NewUserRepo(database *databasePool.Database) *UserRepo {
	return &UserRepo{database}
}

func CreateUser(ctx context.Context, pg *databasePool.Database, user domain.User) error {

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

func GetUsers(ctx context.Context, pg *databasePool.Database) ([]domain.User, error) {

	query := `SELECT * FROM public."user"`
	rows, err := pg.FetchRows(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.User])
}

func (u *UserRepo) GetUserByUsername() (*domain.User, error) {
	return nil, fmt.Errorf("not implemented")
}
