package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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

func Delete(userID int) *sql.DB {
	db := openConnection()
	result, err := db.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	return db
}

func SelectAll() ([]User, *sql.Rows) {
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

	return users, rows
}
