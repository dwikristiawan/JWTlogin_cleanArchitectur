package user

import (
	"JWTLogin/internal/middleware"
	"JWTLogin/internal/model"
	"JWTLogin/internal/repository"
	"JWTLogin/internal/rest/response"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService interface {
	Sign(e echo.Context, user model.Users) (response.Response, error)
	Login(e echo.Context, username string, password string) (response.Response, error)
	RefreshToken(e echo.Context, refreshToken string) (response.Response, error)
}
type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) Sign(e echo.Context, user *model.Users) (response.Response, error) {
	userData, _ := u.repo.FindUserByUsename(e, user.Username)

	if userData.Username != "" {
		return response.Response{StatusCode: 400, Status: "fail", Data: response.Masage{Masage: "username hass bee used"}}, nil
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Response{}, err
	}
	user.Password = string(hashed)
	_, err = u.repo.Create(e, *user)
	if err != nil {
		return response.Response{}, err
	}
	return response.Response{StatusCode: 201, Status: "success", Data: response.Masage{Masage: "sign success"}}, nil

}
func (u UserService) Login(e echo.Context, username string, password string) (response.Response, error) {

	userData, err := u.repo.FindUserByUsename(e, username)
	if err != nil {

		return response.Response{}, err
	}

	if userData.Username == "" {
		return response.Response{StatusCode: http.StatusBadRequest, Status: "failed", Data: response.Masage{Masage: "user not found"}}, err

	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return response.Response{StatusCode: http.StatusBadRequest, Status: "failed", Data: response.Masage{Masage: "wrong password"}}, nil
	}
	token, err := middleware.CreateTokens(userData)
	if err != nil {
		return response.Response{}, err
	}
	return response.Response{StatusCode: http.StatusOK, Status: "success", Data: token}, nil
}

func (u UserService) RefreshToken(e echo.Context, refreshToken string) (response.Response, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) { return middleware.RefreshSecretKey, nil })
	if err != nil || !token.Valid {
		return response.Response{StatusCode: http.StatusUnauthorized, Status: "failed", Data: response.Masage{Masage: "invalid refresh token"}}, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return response.Response{StatusCode: http.StatusUnauthorized, Status: "failed", Data: response.Masage{Masage: "invalid refresh token"}}, nil
	}
	user := model.Users{
		ID:       int(claims["id"].(float64)),
		Username: claims["username"].(string),
		Role:     claims["role"].(string),
	}

	newToken, err := middleware.CreateTokens(user)
	if err != nil {
		return response.Response{}, err
	}

	return response.Response{StatusCode: http.StatusAccepted, Status: "success", Data: newToken}, nil
}
