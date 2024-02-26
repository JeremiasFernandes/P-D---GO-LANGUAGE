package converter

import (
	"github.com/JeremiasFernandes/PeDGo/src/model"
	"github.com/JeremiasFernandes/PeDGo/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
