-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    email varchar(255) UNIQUE,
    password_hash varchar(255),
    token text,
    token_last_updated_at timestamp,
    is_admin boolean DEFAULT FALSE,
    is_premium boolean DEFAULT FALSE,
    created_at timestamp DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;
