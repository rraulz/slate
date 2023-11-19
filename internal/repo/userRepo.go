package repo

import (
	"context"
	"slate/internal/databasePool"
	"slate/internal/domain"

	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	database *databasePool.Database
	ctx      context.Context
}

func NewUserRepo(database *databasePool.Database, ctx context.Context) *UserRepo {
	return &UserRepo{
		database: database,
		ctx:      ctx,
	}
}

func (userRepo *UserRepo) CreateUser(user domain.User) error {

	query := `INSERT INTO "user" (id, username, password, email) VALUES (@id, @username, @password, @email)`
	args := pgx.NamedArgs{
		"id":       user.Id,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
	}

	err := userRepo.database.InsertRow(userRepo.ctx, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (userRepo *UserRepo) GetUsers() ([]domain.User, error) {

	query := `SELECT * FROM public."user"`
	rows, err := userRepo.database.FetchRows(userRepo.ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return pgx.CollectRows(rows, pgx.RowToStructByName[domain.User])
}

func (userRepo *UserRepo) GetUserByUsername(username string) (*domain.User, error) {

	query := `SELECT * FROM public."user" WHERE username = $1`
	row := userRepo.database.QueryRow(userRepo.ctx, query, username)

	user := &domain.User{}
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
