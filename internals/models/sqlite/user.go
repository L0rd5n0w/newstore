package sqlite

import (
	"database/sql"

	//"github.com/mattn/go-sqlite3"
	//"golang.org/x/crypto/bcrypt"
)


type UserModel struct {
	DB	*sql.DB
}

func(bm *UserModel) Insert()