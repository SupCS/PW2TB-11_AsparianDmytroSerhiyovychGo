package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"go2/handlers"
)

// Головна сторінка
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Помилка завантаження сторінки", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	// Налаштування маршрутів
	http.HandleFunc("/", homePage)
	http.HandleFunc("/formula1", handlers.Formula1Handler)
	http.HandleFunc("/formula2", handlers.Formula2Handler)
	http.HandleFunc("/combined", handlers.CombinedHandler)

	// Підключення стилів
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Запуск сервера
	port := ":8080"
	fmt.Println("Сервер запущено на http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
