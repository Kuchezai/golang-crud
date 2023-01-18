package routes

import (
	"CRUD/internal/app/controllers"
	"net/http"
)

func Init() {
	http.HandleFunc("/users", controllers.SelectAll)
	http.HandleFunc("/user/create", controllers.Create)
	http.HandleFunc("/user/update", controllers.Update)
	http.HandleFunc("/user/delete", controllers.Delete)
	http.HandleFunc("/user", controllers.Select)

	http.HandleFunc("/", controllers.HelloServer)

	http.ListenAndServe(":80", nil)
}
