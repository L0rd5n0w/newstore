package sqlite

import (
	"database/sql"
	"log"
	//"github.com/mattn/go-sqlite3"

	"golang.org/x/crypto/bcrypt"
)


type UserModel struct {
	DB	*sql.DB
}

func(bm *UserModel) Insert(id, firstname, lastname, email, password string ) error {
	stmt := `INSERT INTO user (id, firstname, lastname, email, hashed_password, created)
	VALUES (?, ?, ?, ?, ?, datetime("now"))`

	hashed_password, err :=	bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Print(err)
	}

	_, err = bm.DB.Exec(stmt, id, firstname, lastname, email, string(hashed_password))
	if err != nil {
		log.Fatal(err)
	}

	return nil
}