package main

import ( 
	"net/http"	
	"html/template"
)

var templ = template.Must(template.ParseFiles("index.html"))
func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.html",nil)
}

func main() {
	mux := NewServeMux();
	
}
