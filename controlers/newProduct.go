package controlers

import (
	"net/http"
	"strconv"
	"text/template"
	"web_com_golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

// Funcao NewProduct que chama a pagina de novo produto
func NewProduct(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)

}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		name := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		price, errPrice := strconv.ParseFloat(r.FormValue("preco"), 64)
		qtd, errQtd := strconv.Atoi(r.FormValue("quantidade"))

		if errPrice != nil {
			panic(errPrice.Error())
		}
		if errQtd != nil {
			panic(errQtd.Error())
		}

		models.InsertProduct(name, price, descricao, qtd)

	}

	http.Redirect(w, r, "/", 301)
}
