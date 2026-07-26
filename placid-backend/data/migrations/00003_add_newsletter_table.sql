-- +goose Up
CREATE TABLE IF NOT EXISTS newsletter (
    id serial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    created_at timestamp DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS newsletter;
