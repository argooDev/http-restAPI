package main

import (
	"log"

	"github.com/argooDev/http-restAPI/internal/app/apiserver"
)

func main() {
	//Создаем apiserver
	server := apiserver.New()

	//Попытка запуска сервера, если есть ошибка, то выходим через log.Fatal
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
