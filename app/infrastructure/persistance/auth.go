package persistance

import (
	"fmt"

	"github.com/dmskdlghs213/go_authAPI/app/domain/model"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	DB *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		DB: db,
	}
}

func (ar *AuthRepository) CreateStore(storeName string, storeEmail string, storePhoneNumber string) error {
	ins := `
		INSERT INTO store (store_name, store_email, store_phone_number, encrypted_password) VALUES (:store_name, :store_email, :store_phone_number,:encrypted_password)
	`
	_, err := ar.DB.NamedExec(
		ins,
		map[string]interface{}{
			"store_name":         storeName,
			"store_email":        storeEmail,
			"store_phone_number": storePhoneNumber,
			"encrypted_password": "",
		},
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (ar *AuthRepository) FindByStore(storeName string, storeEmail string) (*model.StoreDetail, error) {
	query := `
		SELECT store_id,store_name, store_email,store_phone_number FROM store WHERE store_name = ? AND store_email = ?
	`
	var store model.StoreDetail
	err := ar.DB.Get(&store, query, storeName, storeEmail)
	if err != nil {
		return nil, err
	}

	return &store, nil
}
