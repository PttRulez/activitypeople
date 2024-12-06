-- +goose Up
-- +goose StatementBegin
create table if not exists steps (
	id SERIAL PRIMARY KEY,
	date date NOT NULL,
  steps INT NOT NULL,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE cascade 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table steps;
-- +goose StatementEnd
