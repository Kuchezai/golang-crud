package auth

import (
	logger "CRUD/internal/app/logs"
	"CRUD/internal/app/models"
	"CRUD/internal/auth/services"
	"fmt"
	http "net/http"
)

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
			tokenString := services.CreateTokenWithRole(u.Role)
			services.SetTokenInCookies(w, tokenString)
			http.Redirect(w, r, "/"+u.Role, 301)
			return
		}
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}

func Verification(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	email := query.Get("email")
	sign := query.Get("sign")
	fmt.Println(email, sign)
	if sign == services.GetMD5Hash(email) {
		models.Verification(email)
		http.Redirect(w, r, "/login", 301)
		return
	}
	fmt.Println(email, sign)
	http.Redirect(w, r, "/404", 301)
}
