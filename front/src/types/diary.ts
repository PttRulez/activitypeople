import { ActivityResponse } from "./activity";
import { MealResponse } from "./food";

export type DiaryResponse = {
  activities: ActivityResponse[];
  calories: number;
  caloriesBurned: number;
  caloriesConsumed: number;
  date: string;
  meals: MealResponse[];
  weight: number;
};
