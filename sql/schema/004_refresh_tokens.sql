-- Refresh Tokens
-- This script creates the refresh_tokens table
-- +goose Up
-- goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" up
-- goose postgres "postgres://postgres:postgres@localhost:5432/chirpy" down
-- sqlc generate
-- This migration creates the refresh_tokens table with the necessary fields.
CREATE TABLE IF NOT EXISTS refresh_tokens (
    token TEXT PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    expires_at TIMESTAMP,
    revoked_at TIMESTAMP
);
-- +goose Down  
-- This migration drops the refresh_tokens table.
DROP TABLE IF EXISTS refresh_tokens;



