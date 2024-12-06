-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN calories_per_100_steps INT NOT NULL DEFAULT 5;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN calories_per_100_steps;
-- +goose StatementEnd
