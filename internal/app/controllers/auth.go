package controllers

import (
	logger "CRUD/internal/app/logs"
	"html/template"
	"net/http"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/app/view/layouts/header.gohtml", "internal/app/view/auth/login.gohtml", "internal/app/view/layouts/footer.gohtml")
	if err != nil {
		logger.Error.Println("Error in parsing html-page!")
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "login", nil)
}
