package contracts

type CreateFoodRequest struct {
	Name     string `json:"name" validate:"required"`
	Calories int    `json:"calories" validate:"required,number"`
	Carbs    *int   `json:"carbs" validate:"required,number"`
	Fat      *int   `json:"fat" validate:"required"`
	Protein  *int   `json:"protein" validate:"required"`
}

type FoodResponse struct {
	Name     string `json:"name"`
	Calories int    `json:"calories"`
	Carbs    int    `json:"carbs"`
	Id       int    `json:"id"`
	Fat      int    `json:"fat"`
	Protein  int    `json:"protein"`
}
