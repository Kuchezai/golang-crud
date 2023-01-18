package models

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "gopkg.in/yaml.v2"
)

type DBConnection struct {
	user     string `yaml:"user"`
	password string `yaml:"password"`
	dbname   string `yaml:"dbname"`
	sslmode  string `yaml:"sslmode"`
}

type User struct {
	Id                       int
	Login, Name, Pass, Email string
}

func openConnection() *sql.DB {
	connStr := "user=postgres password=root dbname=CRUDapi sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func Create(login, name, pass, email string) *sql.DB {
	db := openConnection()
	_, err := db.Exec("INSERT INTO users (login, name, pass, email) VALUES ($1, $2, $3, $4)", login, name, pass, email)
	if err != nil {
		panic(err)
	}
	return db
}

func SelectAll() ([]User, *sql.DB) {
	db := openConnection()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	users := []User{}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Login, &u.Name, &u.Pass, &u.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, u)
	}

	return users, db
}

func Select(userID int) (User, *sql.DB) {
	db := openConnection()
	user := User{}
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", userID).Scan(&user.Id, &user.Login, &user.Name, &user.Pass, &user.Email)
	if err != nil {
		panic(err)
	}
	return user, db
}

func Update(userID, login, name, pass, email string) *sql.DB {
	db := openConnection()
	_, err := db.Query("UPDATE users SET login = $2, name = $3, pass = $4, email = $5 WHERE id = $1;", userID, login, name, pass, email)
	if err != nil {
		panic(err)
	}
	return db
}

func Delete(userID int) *sql.DB {
	db := openConnection()
	_, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		panic(err)
	}
	return db
}
