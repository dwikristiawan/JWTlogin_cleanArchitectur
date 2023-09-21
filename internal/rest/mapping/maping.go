package mapping

import (
	"JWTLogin/internal/model"
	"JWTLogin/internal/rest/request"
)

func SignDtoToUser(signDto *request.SignRequest) *model.Users {
	return &model.Users{
		Username: signDto.Username,
		Password: signDto.Password,
		Role:     signDto.Role,
	}
}
