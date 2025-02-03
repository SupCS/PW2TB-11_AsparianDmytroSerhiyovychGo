package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Функція для округлення чисел
func roundFloat(value float64, precision int) string {
	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, value)
}

// Обробник для першого калькулятора
func Formula1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		Q_i, _ := strconv.ParseFloat(r.FormValue("Q_i"), 64)
		B_i, _ := strconv.ParseFloat(r.FormValue("B_i"), 64)
		k_j, _ := strconv.ParseFloat(r.FormValue("k_j"), 64)

		// Розрахунок валового викиду
		E_j := 1e-6 * k_j * Q_i * B_i

		// Відправляємо дані у шаблон
		tmpl, _ := template.ParseFiles("templates/formula1.html")
		tmpl.Execute(w, map[string]string{
			"E_j": roundFloat(E_j, 2),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/formula1.html")
	tmpl.Execute(w, nil)
}
