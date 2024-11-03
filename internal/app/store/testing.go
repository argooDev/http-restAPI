package store

import (
	"fmt"
	"strings"
	"testing"
)

// возвращает тестовый store, который будет сконфигурирован определенным образом
// также функцию, которая позволит очищать заполненные таблицы, чтобы следующие тесты работали с пустой БД
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {

	// Показывает тестам, что этот метод тестовый, его не нужно нигде учитывать и тестировать
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL

	s := New(config) // Создаем хранилище и передаем в него конфиг
	// Попытка открыть хранилище и подключиться к БД, если не получилось - фейлим тесты фаталом
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	// return store and func
	return s, func(tables ...string) {
		// Проверка на наличие таблиц
		if len(tables) > 0 {
			// Соединяемся с БД и выполняем truncate
			if _, err := s.db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ","))); err != nil {
				t.Fatal(err)
			}

		}
		// Закрываем хранилище
		s.Close()
	}
}
