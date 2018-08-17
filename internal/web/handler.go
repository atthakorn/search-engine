package web

import (
	"net/http"
	"html/template"
	"github.com/atthakorn/web-scraper/internal/search"
)


type ViewModel struct {
	Q string
	Result *search.Result
}


func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")


	//create template
	tmpl := template.Must(template.ParseFiles("template/index.gohtml"))


	//parse query
	q := r.URL.Query().Get("q")
	result := search.Query(q)

	//create view model
	model := &ViewModel{Q :q, Result: result}


	tmpl.Execute(w, model)
}
