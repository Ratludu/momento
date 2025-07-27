-- +goose Up
ALTER TABLE profiles ADD COLUMN current_profile INTEGER DEFAULT 0 NOT NULL;

-- +goose Down
ALTER TABLE profiles DROP COLUMN current_profile;
