package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Поля для apiserver
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// Функция New инициализирует apiserver и возвращает сконфигурированный экземпляр  APIServer struct
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start позволяет запусктать http server, connection to database
func (server *APIServer) Start() error {

	// Конфигурируем логгер, проверяем на ошибки
	if err := server.configureLogger(); err != nil {
		return err
	}

	// Вызов роутера
	server.configureRouter()

	// Если никаких ошибок нет и сервак поднялся - выводит инфо-сообщение
	server.logger.Info("Starting apiserver!!!")

	// Добавляем сюда дефолтный адрес и роутер
	return http.ListenAndServe(server.config.BindAddr, server.router)
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

// Описывает обработку входящих запросов
func (server *APIServer) configureRouter() {

	server.router.HandleFunc("/hello", server.handleHello())
}

// Возвращает интерфейс, это позволяет внутри функции определять какие-то переменные, типы и тд
// Позволяет разгрузить код от захламления
func (server *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello web!!!")
	}
}
