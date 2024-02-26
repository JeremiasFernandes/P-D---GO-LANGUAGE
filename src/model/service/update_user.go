package service

import (
	"github.com/JeremiasFernandes/PeDGo/src/configuration/logger"
	"github.com/JeremiasFernandes/PeDGo/src/configuration/rest_err"
	"github.com/JeremiasFernandes/PeDGo/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Iniciando updateUser model.",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Erro ao chamar o repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"updateUser service executedo com successo",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
