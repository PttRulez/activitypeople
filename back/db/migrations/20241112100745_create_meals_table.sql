-- +goose Up
-- +goose StatementBegin
create table if not exists meals (
	id SERIAL PRIMARY KEY,
	calories INT NOT NULL,
	date date NOT NULL,
	name VARCHAR(255) NOT NULL,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE cascade 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table meals;
-- +goose StatementEnd
