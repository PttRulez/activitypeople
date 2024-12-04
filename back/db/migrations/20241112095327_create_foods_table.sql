-- +goose Up
-- +goose StatementBegin
create table if not exists foods (
  carbs INT NOT NULL,
  calories_per_100 INT NOT NULL,
  created_by_admin BOOL NOT NULL,
  fat INT NOT NULL,
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
  protein INT NOT NULL,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE cascade 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table foods;
-- +goose StatementEnd
