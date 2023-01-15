package controllers

import (
	"CRUD/models"
	"fmt"
	"net/http"
	"strconv"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	db := models.Delete(1)
	defer db.Close()
	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		panic(err)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := models.Delete(4)
	defer db.Close()
	_, err := fmt.Fprintf(w, "Пользователь успешно удален!")
	if err != nil {
		panic(err)
	}
}

func SelectUsers(w http.ResponseWriter, r *http.Request) {
	users, db := models.SelectAll()
	defer db.Close()
	for _, u := range users {
		_, err := fmt.Fprintf(w, "Пользователь: %s, %s, %s, %s, %s \n", strconv.Itoa(u.Id), u.Name, u.Email, u.Login, u.Pass)
		if err != nil {
			panic(err)
		}
	}
}
