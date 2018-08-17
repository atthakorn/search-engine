package web

import "html/template"

var funcMap template.FuncMap



func init() {
	funcMap = template.FuncMap{"noescape": noescape}
}



func noescape(str string) template.HTML {
	return template.HTML(str)
}