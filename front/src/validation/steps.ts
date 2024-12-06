import { z } from "zod";

export const CreateStepsSchema = z.object({
  date: z.string(),
  steps: z.number({
    message: "Введите кол-во шагов",
  }),
});

export const UpdateStepsSchema = CreateStepsSchema.partial().extend({
  id: z.number(),
});

export type CreateStepsData = z.infer<typeof CreateStepsSchema>;
export type UpdateStepsData = z.infer<typeof UpdateStepsSchema>;
