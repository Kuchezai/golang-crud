package controllers

import (
	"CRUD/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		panic(err)
	}

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var userID, err = strconv.Atoi(strings.Replace(r.URL.Path[1:], "user/delete/", "", -1))
	if err != nil {
		panic(err)
	}
	db := models.Delete(userID)
	defer db.Close()
	fmt.Fprintf(w, "Пользователь успешно удален!")
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
