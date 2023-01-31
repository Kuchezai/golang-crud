package models

import (
	config "CRUD/internal/app"
	logger "CRUD/internal/app/logs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "gopkg.in/yaml.v2"
)

type User struct {
	Id                                  int
	Login, Name, Pass, Email, Role, Img string
	Verif                               bool
}

func openConnection() *sql.DB {
	configuration := config.GetConfig()
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		(*configuration).DB.User, (*configuration).DB.Password, (*configuration).DB.DBbname, (*configuration).DB.SSLmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error.Println("Open SQL connection:", err)
		panic(err)
	}

	return db
}

func Create(login, name, pass, email, img, role string, verif bool) *sql.DB {
	db := openConnection()
	_, err := db.Exec("INSERT INTO users (login, name, pass, email, img, role, verif) VALUES ($1, $2, $3, $4, $5, $6, $7)", login, name, pass, email, img, role, verif)
	if err != nil {
		logger.Error.Println("Create operation SQL :", err)
		panic(err)
	}
	return db
}

func SelectAll() ([]User, *sql.DB) {
	db := openConnection()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		logger.Error.Println("Select operation SQL :", err)
		panic(err)
	}
	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Login, &u.Name, &u.Pass, &u.Email, &u.Img, &u.Role, &u.Verif)
		if err != nil {
			logger.Error.Println("Incorrect data in SQL :", err)
			panic(err)
		}
		users = append(users, u)
	}

	return users, db
}

func Select(userID int) (User, *sql.DB) {
	db := openConnection()
	user := User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", userID).Scan(&user.Id, &user.Login, &user.Name, &user.Pass, &user.Email, &user.Img, &user.Role, &user.Verif)
	if err != nil {
		logger.Error.Println("Select operation SQL :", err)
		panic(err)
	}
	return user, db
}

func Update(userID, login, name, pass, email, img, role string, verif bool) *sql.DB {
	db := openConnection()
	_, err := db.Query("UPDATE users SET login = $2, name = $3, pass = $4, email = $5, img = $6, role = $7, verif = $8  WHERE id = $1;", userID, login, name, pass, email, img, role, verif)
	if err != nil {
		logger.Error.Println("Update operation SQL :", err)
		panic(err)
	}
	return db
}

func Delete(userID int) *sql.DB {
	db := openConnection()
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		logger.Error.Println("Delete operation SQL :", err)
		panic(err)
	}
	return db
}

func Verification(email string) *sql.DB {
	db := openConnection()
	_, err := db.Exec("UPDATE users SET verif = true WHERE email = $1", email)
	if err != nil {
		logger.Error.Println("Delete operation SQL :", err)
		panic(err)
	}
	return db
}
