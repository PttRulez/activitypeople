-- +goose Up
-- +goose StatementBegin
create table if not exists meals (
	calories INT NOT NULL,
	date date not null,
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE cascade 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table meals;
-- +goose StatementEnd
