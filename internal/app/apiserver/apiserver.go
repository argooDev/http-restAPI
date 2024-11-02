package apiserver

import "github.com/sirupsen/logrus"

// Поля для apiserver
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// Функция New инициализирует apiserver и возвращает сконфигурированный экземпляр  APIServer struct
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start позволяет запусктать http server, connection to database
func (server *APIServer) Start() error {

	// Конфигурируем логгер, проверяем на ошибки
	if err := server.configureLogger(); err != nil {
		return err
	}

	// Если никаких ошибок нет и сервак поднялся - выводит инфо-сообщение
	server.logger.Info("Starting apiserver!!!")
	return nil
}

// Позволяет конфигурировать логгер, может возвращать ошибку из-за неправильного уровня логирования
func (server *APIServer) configureLogger() error {

	// Парсим строку из config.LogLevel
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	// Устанавливаем логгеру соотвествующий уровень
	server.logger.SetLevel(level)
	return nil
}
