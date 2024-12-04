-- +goose Up
-- +goose StatementBegin
create table if not exists users (
	id SERIAL PRIMARY KEY,
  bmr INT NOT NULL default 1800,
  email VARCHAR(255) NOT NULL UNIQUE,
  hashed_password VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  role VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
