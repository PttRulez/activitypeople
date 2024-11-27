export type FoodResponse = {
  name: string;
  calories: number;
  carbs: number;
  fat: number;
  id: number;
  protein: number;
};

export type MealResponse = {
  calories: number;
  date: Date;
  id: number;
  name: string;
  foods: FoodInMealResponse[];
};

type FoodInMealResponse = {
  calories: number;
  name: string;
  weight: number;
};
