package routes

import (
	"net/http"
	"web_com_golang/controlers"
)

func Routes() {
	http.HandleFunc("/", controlers.Index)
}
