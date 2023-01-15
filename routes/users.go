package routes

import (
	"CRUD/controllers"
	"net/http"
)

func Init() {
	http.HandleFunc("/user/delete/", controllers.DeleteUser)
	http.HandleFunc("/users", controllers.SelectUsers)
	http.HandleFunc("/", controllers.HelloServer)

	http.ListenAndServe(":80", nil)
}
