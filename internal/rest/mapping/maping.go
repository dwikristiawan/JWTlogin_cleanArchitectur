package mapping

import (
	"jwtLogin/internal/model"
	"jwtLogin/internal/rest/request"
)

func SignDtoToUser(signDto *request.SignRequest) model.Users {
	return model.Users{
		Username: signDto.Username,
		Password: signDto.Password,
		Role:     signDto.Role,
	}
}
