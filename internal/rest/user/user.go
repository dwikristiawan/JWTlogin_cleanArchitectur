package user

import (
	"jwtLogin/internal/rest/mapping"
	"jwtLogin/internal/rest/request"
	"jwtLogin/internal/rest/response"
	"jwtLogin/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserHandler(e echo.Echo, handler UserHandler) {

	h := handler
	e.POST("/sign", h.SignHandler)
	e.POST("/login", h.LoginHandler)

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
	var responses *response.Response
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}
	if signData.Password == "" || signData.Username == "" || signData.Role == "" {
		responses.Status = "failed"
		responses.StatusCode = http.StatusBadRequest
		responses.Data = response.Masage{Masage: "uncomplete field"}
		return e.JSON(responses.StatusCode, responses)
	}
	responses, err = h.service.Sign(e, mapping.SignDtoToUser(signData))
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	return e.JSON(responses.StatusCode, responses)
}

func (h *UserHandler) LoginHandler(e echo.Context) error {
	var loginData *request.LoginRequest
	var responses *response.Response
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
