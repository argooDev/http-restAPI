package model_test

import (
	"testing"

	"github.com/argooDev/http-restAPI/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t) // Создаем пользователя с помощью хелпера
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
