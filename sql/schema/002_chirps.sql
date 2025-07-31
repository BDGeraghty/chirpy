-- +goose Up
-- goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" up
-- goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" down
-- sqlc generate
-- This migration creates the chirps table with the necessary fields.
CREATE TABLE chirps (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body TEXT NOT NULL
);

-- +goose Down
DROP TABLE chirps;
