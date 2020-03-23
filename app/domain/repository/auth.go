package repository

import "github.com/dmskdlghs213/go_authAPI/app/domain/model"

type AuthAccessor interface {
	CreateStore(storeName string, storeEmail string, storePhoneNumber string) error
	FindByStore(storeName string, storeEmail string) (*model.StoreDetail, error)
}
