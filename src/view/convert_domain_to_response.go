package view

import (
	"github.com/JeremiasFernandes/PeDGo/src/controller/model/response"
	"github.com/JeremiasFernandes/PeDGo/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
