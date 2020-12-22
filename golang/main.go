package main

import (
	"html/template"
    "net/http"
)

func main() {
   http.HandleFunc("/", index)
    http.ListenAndServe(":8000", nil) 
}
func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	data := map[string]string{
		"CodeEducation": greeting("Code.education Rocks!!!"), 
		// A exibição dessa string deve ser baseada no retorno de
		// uma função chamada "greeting". Essa função receberá a
		// string como parâmetro e a retornará entre as tags <b></b>.
	}
	w.WriteHeader(http.StatusOK)
	tpl.Execute(w, data)
}
func greeting(str string)string{
	return str 
}

