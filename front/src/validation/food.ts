import { z } from "zod";

export const CreateFoodSchema = z.object({
  name: z.string(),
  calories: z.number({
    message: "Введите калорийность",
  }),
  carbs: z.number({
    message: "Введите кол-во углеводов",
  }),
  fat: z.number({
    message: "Введите кол-во жиров",
  }),
  protein: z.number({
    message: "Введите кол-во белков",
  }),
});

export const UpdateFoodSchema = CreateFoodSchema.partial().extend({
  id: z.number(),
});

export type CreateFoodData = z.infer<typeof CreateFoodSchema>;
export type UpdateFoodData = z.infer<typeof UpdateFoodSchema>;
