package web

import (
	"net/http"
	"html/template"
	"github.com/atthakorn/search-engine/internal/search"
)

type ViewModel struct {
	Q      string
	Result *search.Result
}




func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")


	//create anonymous template
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("template/index.gohtml"))

	//parse query
	q := r.URL.Query().Get("q")
	result := search.Query(q)

	//create view model
	model := &ViewModel{Q: q, Result: result}

	//execute template
	tmpl.ExecuteTemplate(w, "index.gohtml", model)
}
