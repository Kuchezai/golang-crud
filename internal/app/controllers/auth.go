package controllers

import (
	logger "CRUD/internal/app/logs"
	"CRUD/internal/app/models"
	"CRUD/internal/app/service"
	"html/template"
	http "net/http"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/app/view/layouts/header.gohtml", "internal/app/view/auth/login.gohtml", "internal/app/view/layouts/footer.gohtml")
	if err != nil {
		logger.Error.Println("Error in parsing html-page!")
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "login", nil)
}

func ShowAdmin(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/app/view/layouts/header.gohtml", "internal/app/view/users/admin.gohtml", "internal/app/view/layouts/footer.gohtml")
	if err != nil {
		logger.Error.Println("Error in parsing html-page!")
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "login", nil)
}

func ShowManager(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/app/view/layouts/header.gohtml", "internal/app/view/users/manager.gohtml", "internal/app/view/layouts/footer.gohtml")
	if err != nil {
		logger.Error.Println("Error in parsing html-page!")
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "login", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		logger.Error.Println("Error in parsing form")
		panic(err)
	}
	login := r.FormValue("login")
	pass := r.FormValue("pass")

	users, db := models.SelectAll()
	defer db.Close()
	for _, u := range users {
		if login == u.Login && pass == u.Pass {
			tokenString := service.CreateTokenWithRole(u.Role)
			service.SetTokenInCookies(w, tokenString)
			http.Redirect(w, r, "/"+u.Role, 301)
			return
		}
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
