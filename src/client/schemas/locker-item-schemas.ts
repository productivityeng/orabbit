import { z } from "zod";

export const LockItemFormSchema = z.object({
  reason: z.string().min(10, {
    message: "The reason must be at least 10 characters",
  }),
});
