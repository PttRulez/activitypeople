import { z } from "zod";

const FoodInMealSchema = z.object({
  calories: z.number(),
  caloriesPer100: z.number(),
  id: z.number(),
  name: z.string(),
  weight: z.number(),
});

export const CreateMealSchema = z.object({
  calories: z.number(),
  date: z.string(),
  foods: z.array(FoodInMealSchema),
  name: z.string(),
});

export const UpdateMealSchema = CreateMealSchema.partial().extend({
  id: z.number(),
});

export type CreateMealData = z.infer<typeof CreateMealSchema>;
export type FoodInMealData = z.infer<typeof FoodInMealSchema>;
export type UpdateMealData = z.infer<typeof UpdateMealSchema>;
