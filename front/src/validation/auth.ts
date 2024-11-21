import { z } from "zod";

export const LoginSchema = z.object({
  email: z.string().email(),
  password: z.string(),
});

export type LoginData = z.infer<typeof LoginSchema>;

export const RegisterSchema = LoginSchema.extend({
  confirmPassword: z.string(),
  name: z.string(),
});

export type RegisterData = z.infer<typeof RegisterSchema>;
