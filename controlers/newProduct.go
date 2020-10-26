package controlers

import (
	"net/http"
	"strconv"
	"text/template"
	"web_com_golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	var products = []models.Produto{}

	if models.DataBase == "FILE" {
		products = models.GetAllProductsFile()
	} else {
		products = models.GetAllProducts()
	}

	temp.ExecuteTemplate(w, "Index", products)

}

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

		if models.DataBase == "FILE" {
			models.InsertProductFile(name, price, descricao, qtd)
		} else {
			models.InsertProduct(name, price, descricao, qtd)
		}

	}
	models.SaveProducts()
	http.Redirect(w, r, "/", 301)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	if models.DataBase == "FILE" {
		id, errId := strconv.Atoi(idProduct)
		if errId != nil {
			panic(errId.Error())
		}
		models.DeleteProductFile(id)
	} else {
		models.DeleteProduct(idProduct)
	}
	models.SaveProducts()
	http.Redirect(w, r, "/", 301)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	var product = models.Produto{}

	if models.DataBase == "FILE" {
		id, errId := strconv.Atoi(idProduct)
		if errId != nil {
			panic(errId.Error())
		}
		product = models.GetProductsByIdFile(id)
	} else {
		product = models.GetProductsById(idProduct)
	}

	temp.ExecuteTemplate(w, "New", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
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

		if models.DataBase == "FILE" {
			id, errId := strconv.Atoi(id)
			if errId != nil {
				panic(errId.Error())
			}
			models.UpdateProductFile(id, name, price, descricao, qtd)
		} else {
			models.UpdateProduct(id, name, price, descricao, qtd)
		}

	}
	models.SaveProducts()
	http.Redirect(w, r, "/", 301)

}
