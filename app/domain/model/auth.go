package model

type CreateStore struct {
	StoreName         string `json:"store_name" db:"store_name"`
	StoreEmail        string `json:"store_email" db:"store_email"`
	StorePhoneNumber  string `json:"store_phone_number" db:"store_phone_number"`
	EncryptedPassword string `json:"encrypted_password" db:"encrypted_password"`
}

type StoreDetail struct {
	StoreID          uint   `json:"store_id" db:"store_id"`
	StoreName        string `json:"store_name" db:"store_name"`
	StoreEmail       string `json:"store_email" db:"store_email"`
	StorePhoneNumber string `json:"store_phone_number" db:"store_phone_number"`
}
