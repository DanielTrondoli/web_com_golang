package main

import (
	"html/template"
	"net/http"

	"github.com/DanielTrondoli/web_com_golang/models/produtos"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	products := produtos.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)

}
