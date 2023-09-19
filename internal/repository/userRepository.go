package repository

import (
	"jwtLogin/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type userRepository struct {
	DB *sqlx.DB
}
type UserRepository interface {
	Criete(e echo.Context, userData model.Users) error
	FindUserByUsename(e echo.Context, username string) (model.Users, error)
}

func NewRepository(db *sqlx.DB) *userRepository {
	return &userRepository{DB: db}
}
func (repo userRepository) Criete(e echo.Context, userData model.Users) (model.Users, error) {

	query := `INSERT INTO users(username,role,password)
		VALUES ($1,$2,$3)`
	_, err := repo.DB.NamedExec(query, userData)
	if err != nil {
		return model.Users{}, err
	}
	return model.Users{}, err
}
func (repo userRepository) FindUserByUsename(e echo.Context, username string) (model.Users, error) {
	var userData model.Users
	query := `SELECT * FROM users where username=?`
	err := repo.DB.Select(&userData, query, username)
	if err != nil {
		return model.Users{}, err
	}
	return userData, nil
}
