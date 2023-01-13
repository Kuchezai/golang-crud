package routes

import (
	"CRUD/controllers"
	"net/http"
)

func Handle() {
	http.HandleFunc("/", controllers.HelloServer)
	http.ListenAndServe(":80", nil)
}
