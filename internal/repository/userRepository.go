package repository

import (
	"JWTLogin/internal/model"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type UserRepository struct {
	DB *sqlx.DB
}
type userRepository interface {
	Create(e echo.Context, userData model.Users) error
	FindUserByUsename(e echo.Context, username string) (model.Users, error)
}

func NewRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}
func (repo UserRepository) Create(e echo.Context, userData model.Users) (model.Users, error) {
	var data model.Users
	query := `INSERT INTO users(username,role,password)
		VALUES ($1,$2,$3)`
	err := repo.DB.QueryRowxContext(context.Background(), query, userData.Username, userData.Role, userData.Password).StructScan(&data)
	fmt.Println(err.Error())
	if err != nil && err.Error() != "sql: no rows in result set" {
		return model.Users{}, err
	}
	return data, nil
}
func (repo UserRepository) FindUserByUsename(e echo.Context, username string) (model.Users, error) {
	var userData model.Users
	query := `SELECT * FROM users where username=$1`
	err := repo.DB.QueryRowxContext(context.Background(), query, username).StructScan(&userData)
	if err != nil && err.Error() != "sql: no rows in result set" {
		fmt.Println(err.Error())
		return model.Users{}, err
	}
	return userData, nil
}
