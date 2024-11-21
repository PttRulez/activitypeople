-- +goose Up
-- +goose StatementBegin
create table if not exists activities (
	id SERIAL PRIMARY KEY,
  calories INT NOT NULL,
  description TEXT,
  distance INT NOT NULL,
  date date NOT NULL,
  elevate INT NOT NULL,
  heartrate INT NOT NULL,
  name VARCHAR(255) NOT NULL,
  pace INT NOT NULL,
  pace_string VARCHAR(255) NOT NULL,
	source VARCHAR(255) NOT NULL,
  source_id BIGINT NOT NULL,
	sport_type VARCHAR(255) NOT NULL,
  total_time INT NOT NULL,
	user_id INT NOT NULL REFERENCES users(id) ON DELETE cascade 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table activities;
-- +goose StatementEnd
