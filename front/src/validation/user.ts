import { z } from "zod";

export const SaveSettingsSchema = z.object({
  bmr: z.number({
    message: "Введите ваш базовый обмен",
  }),
  caloriesPer100Steps: z.number({
    message: "Введите кол-во калорий на 100 шагов",
  }),
});

export const UpdateSettingsSchema = SaveSettingsSchema.partial().extend({
  id: z.number(),
});

export type SaveSettingsData = z.infer<typeof SaveSettingsSchema>;
export type UpdateSettingsData = z.infer<typeof UpdateSettingsSchema>;
