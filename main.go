package main

import (
	"net/http"
	"web_com_golang/routes"
)

func main() {

	routes.Routes()
	println("aplicacao iniciada na porta 3000")
	http.ListenAndServe(":3000", nil)

}
