package model

// User соответствует схеме в Бд
type User struct {
	ID                int
	Email             string
	EncryptedPassword string
}
