package controlers

import (
	"net/http"
	"web_com_golang/models"
)

func Index(w http.ResponseWriter, r *http.Request) {

	products := models.GetAllProducts()

	temp.ExecuteTemplate(w, "Index", products)

}
