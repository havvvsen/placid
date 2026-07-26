-- +goose Up
ALTER TABLE IF EXISTS users RENAME COLUMN id TO uuid;

-- +goose Down
ALTER TABLE IF EXISTS users RENAME COLUMN uuid TO id;

