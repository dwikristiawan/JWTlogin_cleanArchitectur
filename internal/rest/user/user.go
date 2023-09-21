package user

import (
	"JWTLogin/app/config"
	"JWTLogin/internal/middleware"
	"JWTLogin/internal/rest/mapping"
	"JWTLogin/internal/rest/request"
	"JWTLogin/internal/rest/response"
	"JWTLogin/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserHandler(e *echo.Echo, rootConfig config.Root, handler UserHandler) {

	h := handler
	e.POST("/sign", h.SignHandler)
	e.POST("/login", h.LoginHandler)
	e.POST("/refresh", h.RefreshTokenHandler)
	e.GET("/user", h.User, middleware.AuthMiddleware("USER"))
	e.GET("/admin", h.Admin, middleware.AuthMiddleware("ADMIN"))

}

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) SignHandler(e echo.Context) error {
	var signData *request.SignRequest
	err := e.Bind(&signData)
	var responses response.Response
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	if signData.Password == "" || signData.Username == "" || signData.Role == "" {
		responses.Status = "failed"
		responses.StatusCode = http.StatusBadRequest
		responses.Data = response.Masage{Masage: "uncomplete field"}
		return e.JSON(responses.StatusCode, responses)
	}
	signUser := mapping.SignDtoToUser(signData)
	responses, err = h.service.Sign(e, signUser)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	return e.JSON(responses.StatusCode, responses)
}

func (h *UserHandler) LoginHandler(e echo.Context) error {
	var loginData *request.LoginRequest
	var responses response.Response

	err := e.Bind(&loginData)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	if loginData.Password == "" || loginData.Username == "" {

		responses.Status = "failed"
		responses.StatusCode = http.StatusBadRequest
		responses.Data = response.Masage{Masage: "uncomplete field"}
		return e.JSON(responses.StatusCode, responses)
	}
	responses, err = h.service.Login(e, loginData.Username, loginData.Password)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(responses.StatusCode, responses)
}
func (h *UserHandler) RefreshTokenHandler(e echo.Context) error {
	refreshToken := e.Request().PostFormValue("refresh_token")
	if refreshToken == "" {
		return e.JSON(http.StatusBadRequest, "refresh token undetected")
	}
	responses, err := h.service.RefreshToken(e, refreshToken)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	return e.JSON(responses.StatusCode, responses)
}

func (h *UserHandler) User(e echo.Context) error {

	return e.JSON(200, response.Response{StatusCode: 200, Status: "success", Data: response.Masage{Masage: "user"}})
}
func (h *UserHandler) Admin(e echo.Context) error {

	return e.JSON(200, response.Response{StatusCode: 200, Status: "success", Data: response.Masage{Masage: "admin"}})
}
