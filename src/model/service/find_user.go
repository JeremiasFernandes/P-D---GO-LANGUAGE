package service

import (
	"github.com/JeremiasFernandes/PeDGo/src/configuration/logger"
	"github.com/JeremiasFernandes/PeDGo/src/configuration/rest_err"
	"github.com/JeremiasFernandes/PeDGo/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Iniciando findUserByEmail services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
