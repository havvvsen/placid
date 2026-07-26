-- +goose Up
ALTER TABLE IF EXISTS users
    ALTER COLUMN password_hash TYPE text;

-- +goose Down
ALTER TABLE IF EXISTS users
    ALTER COLUMN password_hash TYPE varchar(255);

