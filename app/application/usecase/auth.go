package usecase

import (
	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/dmskdlghs213/go_authAPI/app/domain/service"
	"github.com/dmskdlghs213/go_authAPI/app/infrastructure"
	"github.com/dmskdlghs213/go_authAPI/app/infrastructure/persistance"
)

func CreateUser(name string, email string, encrypted_password string) error {
	authRepositoryAccessor := persistance.NewAuthRepository(infrastructure.DB)
	authServiceAccessor := service.NewAuthService(authRepositoryAccessor)

	user := authServiceAccessor.CreateUser(name, email, encrypted_password)

	if user != nil {
		return user
	}

	return user
}

func FindByAccount(name string, email string) (*model.User, error) {
	authRepositoryAccessor := persistance.NewAuthRepository(infrastructure.DB)
	authServiceAccessor := service.NewAuthService(authRepositoryAccessor)

	account, err := authServiceAccessor.FindByAccount(name, email)
	if err != nil {
		return account, err
	}
	
	return account, err
}
