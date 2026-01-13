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
		username string, email string, password string, role string,
	) (*model.User, error)
}

type PGUserRepository struct {
	driver *pgx.Conn
}

func rowToUser(row pgx.Row) (*model.User, error) {
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewPGUserRepository(driver *pgx.Conn) *PGUserRepository {
	return &PGUserRepository{driver: driver}
}

func (r *PGUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := "SELECT * FROM Users WHERE email = $1"
	row := r.driver.QueryRow(ctx, query, email)
	return rowToUser(row)
}

func (r *PGUserRepository) RegisterUser(
	ctx context.Context,
	username string, email string, password string, role string,
) (*model.User, error) {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO Users (Username, Email, Password, Role) 
		VALUES ($1, $2, $3, $4) RETURNING ID, Username, Email, Password, Role`

	row := r.driver.QueryRow(
		ctx, query,
		username, email, passwordHash, role,
	)

	return rowToUser(row)
}
