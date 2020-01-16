package persistance

import (
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

func (ar *AuthRepository) Insert(name string, email string, encryptedPassword string) error {
	ins := `
		INSERT INTO users (name, email, encrypted_password) VALUES (:name, :email, :encrypted_password)
	`
	var user model.User
	_, err := ar.DB.NamedExec(
		ins,
		map[string]interface{}{
			"name":               user.Name,
			"email":              user.Email,
			"encrypted_password": user.EncryptedPassword,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (ar *AuthRepository) FindByAccount(name string, email string) (*model.User, error) {
	query := `
		SELECT user_id,name, email FROM users WHERE name = ? AND email = ?
	`
	var user model.User
	err := ar.DB.Get(&user, query, name, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
