-- +goose Up
CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

-- +goose Down
DROP TABLE user;


CREATE UNIQUE INDEX user_uc_email ON user (email);