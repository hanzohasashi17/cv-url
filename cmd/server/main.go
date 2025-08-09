package main

import (
	"fmt"
	"log"
	"net/http"
	"cv-url/config"
	"cv-url/internal/report"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/send", report.ReportHandler.Create)

	fmt.Printf(`Сервер запущен, порт %v`, config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, r))
}
