-- +goose Up
CREATE TABLE IF NOT EXISTS tracks (
    id serial PRIMARY KEY NOT NULL,
    name varchar(255) UNIQUE NOT NULL,
    mood varchar(50) NOT NULL,
    audio_url text NOT NULL,
    bg_url text NOT NULL,
    created_at timestamp DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS tracks;

