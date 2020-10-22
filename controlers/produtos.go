package controlers

import (
	"net/http"
	"text/template"
	"web_com_golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	products := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)

}
