package apiserver

// APIServer
type APIServer struct {
	config *Config
}

// Функция New возвращает сконфигурированный экземпляр  APIServer struct
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start позволяет запусктать http server, connection to database
func (server *APIServer) Start() error {
	return nil
}
