package store_test

import (
	"testing"

	"github.com/argooDev/http-restAPI/internal/app/model"
	"github.com/argooDev/http-restAPI/internal/app/store"
	"github.com/stretchr/testify/assert"
)

// Тесты для метода Create
func TestUserRepository_Create(t *testing.T) {
	// Назнаем хранилище и teardown функцию, которая возвращается в store.TestStore вместе с вызовом этой функции
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users") // Вызываем функцию teardown и очищаем таблицу users

	// Создаем пользователя, передаем модель юзера в метод create
	u, err := s.User().Create(model.TestUser(t))
	// С помощью testify проверяем что нет ошибок и что юзер не nil
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

// Тесты для метода FindByEmail
func TestUserRepository_FindByEmail(t *testing.T) {
	// Назнаем хранилище и teardown функцию, которая возвращается в store.TestStore вместе с вызовом этой функции
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users") // Вызываем функцию teardown и очищаем таблицу users

	// 1 случай: Попытка найти несуществующего пользователя - получаем ошибку
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err) // Проверяем корректность

	u := model.TestUser(t)
	u.Email = email

	// 2 случай: Пользователь есть в БД, попытка прочитать его
	s.User().Create(u)
	// Не ожидаем ошибку, ждем что юзер != nil
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
