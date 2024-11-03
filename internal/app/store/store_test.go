package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// Вызывается один раз в пакете перед всеми тестами
// Хороший способ делать разовые манипуляции, например, читать env переменные значения и тд
func TestMain(m *testing.M) {

	// Назначаем глобальную переменную из env переменной
	databaseURL = os.Getenv("DATABASE_URL")
	// Если не получается, то юзаем дефолтное значение
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}

	// Выходим с правильным кодом
	os.Exit(m.Run())
}
