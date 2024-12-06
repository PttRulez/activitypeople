package contracts

type UserSettingsRequest struct {
	BMR                 int     `json:"bmr" validate:"required,number"`
	CaloriesPer100Steps float64 `json:"caloriesPer100Steps" validate:"required,number"`
}
