package contracts

type CreateFoodRequest struct {
	Name     string `json:"name" validate:"required"`
	Calories int    `json:"calories" validate:"required,int"`
	Carbs    int    `json:"carbs" validate:"required"`
	Fat      int    `json:"fat" validate:"required"`
	Protein  int    `json:"protein" validate:"required"`
	Public   bool   `json:"public" validate:"required"`
}
