package user

import (
	"jwtLogin/internal/middleware"
	"jwtLogin/internal/model"
	"jwtLogin/internal/repository"
	"jwtLogin/internal/rest/response"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Sign(e echo.Context, user model.Users) (*response.Response, error)
	Login(e echo.Context, username string, password string) (*response.Response, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo: repo}
}

func (u userService) Sign(e echo.Context, user *model.Users) (*response.Response, error) {
	userData, err := u.repo.FindUserByUsename(e, user.Username)
	if err != nil {
		return nil, err
	}
	if userData.Username == "" {
		return &response.Response{StatusCode: 400, Status: "fail", Data: response.Masage{Masage: "user not found"}}, nil
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashed)
	_, err = u.repo.Criete(e, *user)
	if err != nil {
		return nil, err
	}
	return &response.Response{StatusCode: 201, Status: "success", Data: response.Masage{Masage: "sign success"}}, nil

}
func (u userService) Login(e echo.Context, username string, password string) (*response.Response, error) {
	userData, err := u.repo.FindUserByUsename(e, username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	token, err := middleware.CreateTokens(userData)
	if err != nil {
		return nil, err
	}
	return &response.Response{StatusCode: http.StatusOK, Status: "succes", Data: token}, nil
}
