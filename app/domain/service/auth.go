package service

import (
	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/dmskdlghs213/go_authAPI/app/domain/repository"
)

type AuthService struct {
	authAccessor repository.AuthAccessor
}

func NewAuthService(ur repository.AuthAccessor) *AuthService {
	return &AuthService{
		authAccessor: ur,
	}
}

func (as *AuthService) CreateStore(storeName string, storeEmail string, storePhoneNumber string) error {

	err := as.authAccessor.CreateStore(storeName, storeEmail, storePhoneNumber)
	if err != nil {
		return err
	}

	return err
}

func (as *AuthService) FindByStore(storeName string, storeEmail string) (*model.StoreDetail, error) {
	store, err := as.authAccessor.FindByStore(storeName, storeEmail)
	if err != nil {
		return nil, err
	}

	return store, err
}
