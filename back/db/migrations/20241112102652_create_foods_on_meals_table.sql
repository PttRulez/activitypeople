-- +goose Up
-- +goose StatementBegin
create table if not exists foods_on_meals (
	calories INT NOT NULL,
	meal_id INT NOT NULL REFERENCES meals(id) ON DELETE cascade,
	name VARCHAR(255) NOT NULL,
	food_id INT NOT NULL REFERENCES foods(id) ON DELETE cascade,
	weight INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table foods_on_meals;
-- +goose StatementEnd
