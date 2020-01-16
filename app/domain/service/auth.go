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

func (as *AuthService) CreateUser(name string, email string, encrypted_password string) error {

	user := as.authAccessor.Insert(name, email, encrypted_password)
	if user != nil {
		return user
	}

	return user
}

func (as *AuthService) FindByAccount(name string, email string) (*model.User, error) {
	account, err := as.authAccessor.FindByAccount(name, email)
	if err != nil {
		return nil, err
	}

	return account, err
}
