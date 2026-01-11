package model

type RoleEnum int

const (
	Student = iota
	Teacher
	Admin
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
	Role     RoleEnum
}
