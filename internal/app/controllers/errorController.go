package controllers

import (
	logger "CRUD/internal/app/logs"
	"html/template"
	"net/http"
)

func Error404(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("internal/app/view/errors/404.gohtml")
	if err != nil {
		logger.Error.Println("Error in parsing html-page!")
		panic(err)
	}

	tmpl.ExecuteTemplate(w, "error", nil)
}
