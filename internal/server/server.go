package server

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"log"
	"net/http"
	"time"
)

type Serv struct {
	loggerServ *log.Logger
	HTTPServer http.Server
}

func Myserver(logger *log.Logger) *Serv {
	router := http.NewServeMux() 
router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
	 handlers.HandleGet(res, req, logger)})
router.HandleFunc("/upload", func(res http.ResponseWriter, req *http.Request){
	handlers.HandlePost(res, req, logger)})
 S := &Serv{
		loggerServ: logger,
		HTTPServer: http.Server{
			Addr:                         ":8080",
			Handler:                      router,
			ReadTimeout:                  5 * time.Second,
			WriteTimeout:                 10 * time.Second,
			IdleTimeout:                  15 * time.Second,
			ErrorLog:                     logger,
					},
	}
	return S
}


