package repository

import "github.com/dmskdlghs213/go_authAPI/app/domain/model"

type AuthAccessor interface {
	Insert(name string, email string, encrypted_password string) error
	FindByAccount(name string, email string) (*model.User,error)
}
