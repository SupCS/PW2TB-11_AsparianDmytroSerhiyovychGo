package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

func Formula2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Q_i, _ := strconv.ParseFloat(r.FormValue("Q_i"), 64)
		A_r, _ := strconv.ParseFloat(r.FormValue("A_r"), 64)
		G_viv, _ := strconv.ParseFloat(r.FormValue("G_viv"), 64)
		efficiency, _ := strconv.ParseFloat(r.FormValue("efficiency"), 64)
		a_viv, _ := strconv.ParseFloat(r.FormValue("a_viv"), 64)
		k_tvs, _ := strconv.ParseFloat(r.FormValue("k_tvs"), 64)

		// Розрахунок показника емісії твердих частинок
		k_tv := (1e6/Q_i)*a_viv*(A_r/(100-G_viv))*(1-efficiency) + k_tvs

		// Відправляємо дані у шаблон
		tmpl, _ := template.ParseFiles("templates/formula2.html")
		tmpl.Execute(w, map[string]string{
			"k_tv": roundFloat(k_tv, 2),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/formula2.html")
	tmpl.Execute(w, nil)
}
