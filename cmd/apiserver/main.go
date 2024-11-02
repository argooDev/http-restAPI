package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/argooDev/http-restAPI/internal/app/apiserver"
)

// Путь к apiserver.toml должен задаваться в качестве флага при запуске бинарника
var (
	configPath string
)

// Позволяет парсить configPath, флаг config-path, значение по умолчанию, описание
func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	// Позволяет распарсить флаги выше и записать в переменные
	flag.Parse()

	//Инициализируем конфиг, чтобы передать его в фукнцию New()
	config := apiserver.NewConfig()

	// Читаем файл apiserver.toml и записываем значения в config, после проверяем на возможные ошибки
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	//Создаем apiserver
	server := apiserver.New(config)

	//Попытка запуска сервера, если есть ошибка, то выходим через log.Fatal
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
