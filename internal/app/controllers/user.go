package controllers

import (
	"CRUD/internal/app/models"
	"fmt"
	"net/http"
	"strconv"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		panic(err)
	}

}

func Create(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	login := query.Get("login")
	name := query.Get("name")
	pass := query.Get("pass")
	email := query.Get("email")

	db := models.Create(login, name, pass, email)
	defer db.Close()

	fmt.Fprintf(w, "Пользователь успешно добавлен!")
}

func SelectAll(w http.ResponseWriter, r *http.Request) {
	users, db := models.SelectAll()
	defer db.Close()
	for _, u := range users {
		_, err := fmt.Fprintf(w, "Пользователь: %s, %s, %s, %s, %s \n", strconv.Itoa(u.Id), u.Name, u.Email, u.Login, u.Pass)
		if err != nil {
			panic(err)
		}
	}
}

func Select(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		panic(err)
	}
	user, db := models.Select(userID)
	defer db.Close()

	_, err = fmt.Fprintf(w, "Запрашиваемый пользователь: %s, %s, %s, %s, %s \n", strconv.Itoa(user.Id), user.Name, user.Email, user.Login, user.Pass)
	if err != nil {
		panic(err)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	login := query.Get("login")
	name := query.Get("name")
	pass := query.Get("pass")
	email := query.Get("email")

	db := models.Update(id, login, name, pass, email)
	defer db.Close()

	fmt.Fprintf(w, "Пользователь успешно добавлен!")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		panic(err)
	}
	db := models.Delete(userID)
	defer db.Close()
	fmt.Fprintf(w, "Пользователь успешно удален!")
}
