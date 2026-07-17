-- +goose Up
CREATE TABLE books (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    descrip TEXT NOT NULL,
    createdAt DATETIME NOT NULL
);

-- +goose Down
DROP TABLE books;
