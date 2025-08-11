package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger := log.New(file, ".....", log.Ldate|log.Ltime|log.Lshortfile)
	Srv := server.Myserver(logger)
	logger.Println("Порт", Srv.HTTPServer.Addr)
    err=Srv.HTTPServer.ListenAndServe() 
    if err != nil {
		logger.Fatal("Ошибка при запуске сервера: ", err)
	}
}
