package sqlite

import (
	"database/sql"

	"github.com/L0rd5n0w/newstore/internals/models"
)

type BooksModel struct {
	DB *sql.DB
}

func(bm *BooksModel) Get(id int) (*models.Books, error) {
	stmt := `SELECT id, title, author, descrip, createdAt FROM books WHERE id = ?`

	row := bm.DB.QueryRow(stmt, id)
	b := &models.Books{}

	err := row.Scan(&b.ID, &b.Author, &b.Description, &b.CreatedAt)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func(bm *BooksModel) All() ([]models.Books, error) {
	stmt := `SELECT id, title, author, descrip, createdAt FROM books ORDER BY id DESC`
	rows, err := bm.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	books := []models.Books{}
	for rows.Next() {
		b := models.Books{}
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Description, &b.CreatedAt)
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
	stmt := `INSERT INTO books (title, author, descrip, createdAt)
	VALUES(?, ?, ?, datetime("now"))`

	_, err := bm.DB.Exec(stmt, title, author, description)
	return err
}

func(bm *BooksModel) Update(author, description string) error {
	stmt := `UPDATE books SET (author, descrip) WHERE id
	VALUES(?, ?, ?)`

	_, err := bm.DB.Exec(stmt, author, description)
	if err != nil {
		return err
	}

	return nil
}