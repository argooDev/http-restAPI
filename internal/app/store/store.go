package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

// Обрабатывает store и возвращает уже сконфигурированный
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Инициализация хранилища, попытка подключения к БД и тд
func (server *Store) Open() error {

	// Подключаемся к db с помощью метода sql.Open, драйвер - postgres, sourceName - DatabaseURL
	db, err := sql.Open("postgres", server.config.DatabaseURL)
	if err != nil {
		return err
	}

	// Проверка соединения с БД с помощью Ping (Он под капотом просто select вызывает)
	if err := db.Ping(); err != nil {
		return err
	}

	server.db = db

	return nil
}

// Отключение от БД и тд
func (server *Store) Close() {
	server.db.Close()
}
