package routes

import (
	"net/http"
	"web_com_golang/controlers"
)

func Routes() {
	http.HandleFunc("/", controlers.Index)
	http.HandleFunc("/new", controlers.NewProduct)
	http.HandleFunc("/insert", controlers.InsertProduct)
	http.HandleFunc("/delete", controlers.DeleteProduct)

}
