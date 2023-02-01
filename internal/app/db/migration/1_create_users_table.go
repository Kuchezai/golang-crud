package main

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table users...")
		_, err := db.Exec(`CREATE TABLE users
		(
			id serial NOT NULL,
			login character varying(50) NOT NULL,
			name character varying(50) NOT NULL,
			pass character varying(50) NOT NULL,
			email character varying(50) NOT NULL,
			img character varying(50) NOT NULL,
			role character varying(50) NOT NULL,
			verif bool NOT NULL,
			PRIMARY KEY (id)
		);`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table users...")
		_, err := db.Exec("DROP TABLE IF EXISTS  users")
		return err
	})
}
