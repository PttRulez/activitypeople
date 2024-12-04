import { z } from "zod";

export const CreateWeightSchema = z.object({
  date: z.string(),
  weight: z.number({
    message: "Введите вес",
  }),
});

export const UpdateWeightSchema = CreateWeightSchema.partial().extend({
  id: z.number(),
});

export type CreateWeightData = z.infer<typeof CreateWeightSchema>;
export type UpdateWeightData = z.infer<typeof UpdateWeightSchema>;
