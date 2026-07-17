package sqlite

import (
	"database/sql"
)

type BooksModel struct {
	DB *sql.DB
}