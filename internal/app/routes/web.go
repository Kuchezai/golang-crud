package routes

import (
	config "CRUD/internal/app"
	"CRUD/internal/app/controllers"
	logger "CRUD/internal/app/logs"
	"CRUD/internal/app/middlewares"
	"CRUD/internal/auth"
	"github.com/gorilla/mux"
	"net/http"
)

func Init() {
	configuration := config.GetConfig()

	m := mux.NewRouter()

	/*----------CRUD GET запросы------------*/
	m.HandleFunc("/users", controllers.SelectAll).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/create", controllers.Create).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/update", controllers.Update).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user/delete", controllers.Delete).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/user", controllers.Select).Host((*configuration).Server.Host).Methods("GET")

	/*----------CRUD POST запросы и GUI------------*/
	m.HandleFunc("/login", controllers.ShowLogin).Host((*configuration).Server.Host).Methods("GET")
	m.HandleFunc("/login", auth.Login).Host((*configuration).Server.Host).Methods("POST")
	m.HandleFunc("/manager", controllers.ShowManager).Host((*configuration).Server.Host).Methods("GET")

	/*----------AUTH AND ADMIN------------*/
	m.Handle("/admin", middlewares.IsAdmin(http.HandlerFunc(controllers.ShowAdmin))).Methods("GET")
	m.HandleFunc("/user/create/ui", controllers.ShowCreate).Host((*configuration).Server.Host).Methods("GET")
	m.Handle("/user/create", middlewares.IsAdmin(http.HandlerFunc(controllers.CreateUser))).Methods("POST")
	m.HandleFunc("/verification", auth.Verification).Host((*configuration).Server.Host).Methods("GET")

	/*----------Errors------------*/
	m.HandleFunc("/404", controllers.Error404).Host((*configuration).Server.Host).Methods("GET")

	logger.Info.Println("Server start!")
	err := http.ListenAndServe((*configuration).Server.Port, m)
	if err != nil {
		logger.Error.Println("Server didn't start!")
		panic(err)
	}
}
