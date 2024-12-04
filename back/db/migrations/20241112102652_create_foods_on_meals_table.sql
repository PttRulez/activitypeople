-- +goose Up
-- +goose StatementBegin
create table if not exists foods_on_meals (
	calories INT NOT NULL,
	calories_per_100 INT NOT NULL,
	meal_id INT NOT NULL REFERENCES meals(id) ON DELETE cascade,
	name VARCHAR(255) NOT NULL,
	weight INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table foods_on_meals;
-- +goose StatementEnd
