package routes

import (
	config "CRUD/internal/app"
	"CRUD/internal/app/controllers"
	logger "CRUD/internal/app/logs"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	configuration := config.GetConfig()

	m := mux.NewRouter()
	m.HandleFunc("/users", controllers.SelectAll).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/create", controllers.Create).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/update", controllers.Update).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/delete", controllers.Delete).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user", controllers.Select).Host((*configuration).Server.Host).Methods("GET")

	m.HandleFunc("/user/create/ui", controllers.ShowUi).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/create", controllers.CreateFromUi).Host((*configuration).Server.Host).Methods("POST")

	logger.Info.Println("Server start!")
	err := http.ListenAndServe((*configuration).Server.Port, m)
	if err != nil {
		logger.Error.Println("Server didn't start!")
		panic(err)
	}
}
