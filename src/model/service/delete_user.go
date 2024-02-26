package service

import (
	"github.com/JeremiasFernandes/PeDGo/src/configuration/logger"
	"github.com/JeremiasFernandes/PeDGo/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string) *rest_err.RestErr {

	logger.Info("Iniciando deleteUser model.",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Erro ao chamar o repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info(
		"deleteUser service executado com successo",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))
	return nil
}
