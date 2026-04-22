package main

import (
	"html/template"
	"log"
	"net/http"

	"dork-otomasyon/internal/handler"
	"dork-otomasyon/internal/repository"
	"dork-otomasyon/internal/service"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	repo := repository.NewDorkRepository("data/dorks.json")
	svc := service.NewDorkService(repo)
	appHandler := handler.NewHandler(svc, tmpl)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", appHandler.Home)
	http.HandleFunc("/generate", appHandler.Generate)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
