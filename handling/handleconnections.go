package handling

import (
	"html/template"
	"net/http"
)

func Handlemain(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./static/template.html")
	t.Execute(w, t)
}
