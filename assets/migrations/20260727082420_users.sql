-- +goose Up
CREATE TABLE user (
    id TEXT PRIMARY KEY, -- UUID stored as text
    firstname VARCHAR(255) NOT NULL,
    lastname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created DATETIME NOT NULL
);

-- +goose Down



CREATE UNIQUE INDEX user_uc_email ON user (email);
