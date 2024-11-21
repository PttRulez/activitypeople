-- +goose Up
-- +goose StatementBegin
create table if not exists foods_on_meals (
	meal_id INT NOT NULL REFERENCES meals(id) ON DELETE cascade,
	food_id INT NOT NULL REFERENCES foods(id) ON DELETE cascade,
	weight int NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table foods_on_meals;
-- +goose StatementEnd
