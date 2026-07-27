package sqlite

import (
	"database/sql"
	"log"
	//"github.com/mattn/go-sqlite3"
	//"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

)


type UserModel struct {
	DB	*sql.DB
}

func(bm *UserModel) Insert(id, firstname, lastname, email string, hashed_password []byte) error {
	stmt := `INSERT INTO user (id, firstname, lastname, email, hashed_password, created)
	VALUES (?, ?, ?, ?, ?, datetime("now"))`

	_, err := bm.DB.Exec(stmt, id, firstname, lastname, email, hashed_password)
	if err != nil {
		log.Fatal(err)
	}

	id = uuid.NewString()

	return nil
}