package controllers

import (
	logger "CRUD/internal/app/logs"
	"CRUD/internal/app/models"
	"CRUD/internal/app/service"
	smtp "CRUD/internal/smtp"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	login := query.Get("login")
	name := query.Get("name")
	pass := query.Get("pass")
	email := query.Get("email")
	img := query.Get("img")
	role := query.Get("role")
	verif := query.Get("verif")
	verifStr, err := strconv.ParseBool(verif)
	if err != nil {
		logger.Error.Println("Parse request ", err)
		panic(err)
	}

	db := models.Create(login, name, pass, email, img, role, verifStr)
	defer db.Close()

	fmt.Fprintf(w, "Пользователь успешно добавлен!")
}

func SelectAll(w http.ResponseWriter, r *http.Request) {
	users, db := models.SelectAll()
	defer db.Close()
	for _, u := range users {
		_, err := fmt.Fprintf(w, "Пользователь: %s, %s, %s, %s, %s, %s, %s \n", strconv.Itoa(u.Id), u.Name, u.Email, u.Login, u.Pass, u.Img, u.Role)
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

	_, err = fmt.Fprintf(w, "Запрашиваемый пользователь: %s, %s, %s, %s, %s, %s, %s \n", strconv.Itoa(user.Id), user.Name, user.Email, user.Login, user.Pass, user.Img, user.Role)
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
	img := query.Get("img")
	role := query.Get("role")
	verif := query.Get("verif")
	verifStr, err := strconv.ParseBool(verif)
	if err != nil {
		logger.Error.Println("Parse request ", err)
		panic(err)
	}

	db := models.Update(id, login, name, pass, email, img, role, verifStr)
	defer db.Close()

	fmt.Fprintf(w, "Пользователь успешно добавлен!")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID, err := strconv.Atoi(query.Get("id"))
	if err != nil {
		logger.Error.Println("Incorrect ID")
		panic(err)
	}
	db := models.Delete(userID)
	defer db.Close()

	fmt.Fprintf(w, "Пользователь успешно удален!")
}

func ShowCreate(w http.ResponseWriter, r *http.Request) {
	users, db := models.SelectAll()
	defer db.Close()
	tmpl, err := template.ParseFiles("internal/app/view/layouts/header.gohtml", "internal/app/view/users/create.gohtml", "internal/app/view/layouts/footer.gohtml")
	if err != nil {
		logger.Error.Println("Error in parsing html-page!")
		panic(err)
	}
	tmpl.ExecuteTemplate(w, "create", users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		logger.Error.Println("Error in parsing form")
		panic(err)
	}

	login := r.FormValue("login")
	name := r.FormValue("name")
	pass := r.FormValue("pass")
	email := r.FormValue("email")
	role := r.FormValue("role")

	file, handler, err := r.FormFile("image")
	img := service.UploadImageFromForm(file, handler)

	db := models.Create(login, name, pass, email, img, role, false)
	defer db.Close()

	mailer := smtp.NewMailer()
	mailer.SendVerificationMail(email)

	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}

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
