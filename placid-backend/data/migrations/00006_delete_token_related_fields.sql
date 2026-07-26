-- +goose Up
ALTER TABLE IF EXISTS users
    DROP COLUMN token_last_updated_at;

ALTER TABLE IF EXISTS users
    DROP COLUMN token;

-- +goose Down
ALTER TABLE IF EXISTS users
    ADD COLUMN token TYPE text;

ALTER TABLE IF EXISTS users
    ADD COLUMN token_last_updated_at TYPE timestamp;
