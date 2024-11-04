package model_test

import (
	"testing"

	"github.com/argooDev/http-restAPI/internal/app/model"
	"github.com/stretchr/testify/assert"
)

// Метод для тестирования валидации
// Содержит в себе табличные тесты
func TestUser_Validate(t *testing.T) {

	// testCases - массив анонимных структур
	testCases := []struct {
		name    string
		u       func() *model.User // Фукнция, которая возвращает юзера, для того, чтобы внутри этой функции можно было менять параметры для тестов
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"

				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""

				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)

				u.Password = "short"
				return u
			},
			isValid: false,
		},
	}

	// Итерирует по testCases и запускает их
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}

		})
	}

}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t) // Создаем пользователя с помощью хелпера
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
