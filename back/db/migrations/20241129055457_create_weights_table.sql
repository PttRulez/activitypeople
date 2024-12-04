-- +goose Up
-- +goose StatementBegin
create table if not exists weights (
	id SERIAL PRIMARY KEY,
	date date NOT NULL,
  weight numeric(5,1) NOT NULL,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE cascade 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table weights;
-- +goose StatementEnd
