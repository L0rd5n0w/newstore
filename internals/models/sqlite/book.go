package sqlite

import (
	"database/sql"

	"github.com/L0rd5n0w/newstore/internals/models"
)

type BooksModel struct {
	DB *sql.DB
}

func(bm *BooksModel) All() ([]models.Books, error) {
	stmt := `SELECT id, title, author, createdAT FROM books ORDER BY id DESC`
	rows, err := bm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	books := []models.Books{}
	for rows.Next() {
		b := models.Books{}
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func(bm *BooksModel) Insert(title, author, description string) error {
	stmt := `INSERT INTO books (title, author, description, createdAT)
	VALUES(?, ?, ?, datetime("now"))`

	_, err := bm.DB.Exec(stmt, title, author, description)
	return err
}