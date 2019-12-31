package service

import "https://github.com/dmskdlghs213/go_authAPI/tree/master/app/domain/model"

type AuthService struct {
	userAccessor repository.UserAccessor
}

func NewAuthService(ua repository.UserAccessor) AuthService {
	return &AuthService{
		userAccessor: ua,
	}
}
