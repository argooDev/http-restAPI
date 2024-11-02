package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"` // Адрес запуска веб сервера
}

// NewConfig - возвращает конфигурацию с дефолтными параметрами
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
