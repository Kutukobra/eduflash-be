package repository

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	RegisterUser(
		ctx context.Context,
		username string, email string, password string,
	) error
}

type PGUserRepository struct {
	driver *pgx.Conn
}

func rowToUser(row pgx.Row) (*model.User, error) {
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewPGUserRepository(driver *pgx.Conn) *PGUserRepository {
	return &PGUserRepository{driver: driver}
}

func (r *PGUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT ID, Username, Email, Password FROM Users WHERE email = $1"
	row := r.driver.QueryRow(ctx, query, email)
	return rowToUser(row)
}

func (r *PGUserRepository) RegisterUser(
	ctx context.Context,
	username string, email string, password string,
) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO Users (Username, Email, Password) VALUES ($1, $2, $3)`

	_, err = r.driver.Exec(
		ctx, query,
		username, email, passwordHash,
	)

	return err
}
