package model

type User struct {
	UserID uint              `json:"user_id" db:"user_id"`
	Name string              `json:"name" db:"name"`
	Email string             `json:"email" db:"email"`
	EncryptedPassword string `json:"encrypted_password" db:"encrypted_password"`
}
