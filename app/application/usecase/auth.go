package usecase

import (
	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/dmskdlghs213/go_authAPI/app/domain/service"
	"github.com/dmskdlghs213/go_authAPI/app/infrastructure"
	"github.com/dmskdlghs213/go_authAPI/app/infrastructure/persistance"
)

func CreateStore(storeName string, storeEmail string, storePhoneNumber string) error {
	authRepositoryAccessor := persistance.NewAuthRepository(infrastructure.DB)
	authServiceAccessor := service.NewAuthService(authRepositoryAccessor)

	err := authServiceAccessor.CreateStore(storeName, storeEmail, storePhoneNumber)

	if err != nil {
		return err
	}

	return err
}

func FindByStore(storeName string, storeEmail string) (*model.StoreDetail, error) {
	authRepositoryAccessor := persistance.NewAuthRepository(infrastructure.DB)
	authServiceAccessor := service.NewAuthService(authRepositoryAccessor)

	store, err := authServiceAccessor.FindByStore(storeName, storeEmail)
	if err != nil {
		return store, err
	}

	return store, err
}
