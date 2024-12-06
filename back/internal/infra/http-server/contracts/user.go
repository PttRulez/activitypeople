package contracts

type UserSettingsRequest struct {
	BMR                 int `json:"bmr" validate:"required,number"`
	CaloriesPer100Steps int `json:"caloriesPer100Steps" validate:"required,number"`
}
