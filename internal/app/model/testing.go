package model

import "testing"

// Позволяет избегать постоянной инициализации этой структуры в тестах
// Этот хелпер просто будет возвращать структурку с валидными данными
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}
