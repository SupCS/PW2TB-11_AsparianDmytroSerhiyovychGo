package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

func CombinedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Q_i, _ := strconv.ParseFloat(r.FormValue("Q_i"), 64)
		A_r, _ := strconv.ParseFloat(r.FormValue("A_r"), 64)
		G_viv, _ := strconv.ParseFloat(r.FormValue("G_viv"), 64)
		efficiency, _ := strconv.ParseFloat(r.FormValue("efficiency"), 64)
		a_viv, _ := strconv.ParseFloat(r.FormValue("a_viv"), 64)
		B_i, _ := strconv.ParseFloat(r.FormValue("B_i"), 64)
		k_tvs, _ := strconv.ParseFloat(r.FormValue("k_tvs"), 64)

		k_tv := (1e6/Q_i)*a_viv*(A_r/(100-G_viv))*(1-efficiency) + k_tvs
		E_j := 1e-6 * k_tv * Q_i * B_i

		tmpl, _ := template.ParseFiles("templates/combined.html")
		tmpl.Execute(w, map[string]string{
			"k_tv": roundFloat(k_tv, 2),
			"E_j":  roundFloat(E_j, 2),
		})
		return
	}

	tmpl, _ := template.ParseFiles("templates/combined.html")
	tmpl.Execute(w, nil)
}
