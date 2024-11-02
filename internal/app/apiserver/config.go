package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"` // Адрес запуска веб сервера
	LogLevel string `toml:"log_level"` // Уровень логирования
}

// NewConfig - возвращает конфигурацию с дефолтными параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080", // Дефолт значение адреса веб сервера
		LogLevel: "debug", // Дефолтный уровень логирования установлен на debug
	}
}
