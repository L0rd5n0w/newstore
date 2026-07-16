-- +goose Up
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    createdAT DATETIME NOT NULL
);



-- +goose Down
DROP TABLE books;
