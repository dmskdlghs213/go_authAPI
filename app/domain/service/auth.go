package service

import "github.com/dmskdlghs213/go_authAPI/app/domain/repository"

type AuthService struct {
	userAccessor repository.UserAccessor
}

func NewAuthService(ua repository.UserAccessor) *AuthService {
	return &AuthService{
		userAccessor: ua,
	}
}
