package repository

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string)
	RegisterUser(
		ctx context.Context,
		username string, email string, password string, role model.RoleEnum,
	)
	LoginUser(
		ctx context.Context,
		email string, password string,
	)
}

type PGUserRepository struct {
	driver *pgx.Conn
}

func rowToUser(row pgx.Row) (*model.User, error) {
	var roleString string
	var user model.User
	err := row.Scan(&user.Username, &user.Email, &user.Password, &roleString)
	if err != nil {
		return nil, err
	}

	switch roleString {
	case "Student":
		user.Role = model.Student
	case "Teacher":
		user.Role = model.Teacher
	case "Admin":
		user.Role = model.Admin
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
	username string, email string, password string, role model.RoleEnum,
) (*model.User, error) {
	var roleString string
	switch role {
	case model.Teacher:
		roleString = "Teacher"
	case model.Student:
		roleString = "Student"
	case model.Admin:
		roleString = "Admin"
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO Users (Username, Email, Password, Role) 
		VALUES ($1, $2, $3) RETURNING ID, Username, Email, Password, Role`

	row := r.driver.QueryRow(
		ctx, query,
		username, passwordHash, roleString,
	)

	return rowToUser(row)
}

func (r *PGUserRepository) LoginUser(ctx context.Context, email string, password string) (*model.User, error) {
	userData, err := r.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	userData.Password = "" // clean so no exposed hash

	return userData, nil
}
