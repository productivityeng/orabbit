import { z } from "zod";

export const CreateRabbitMqQeueueRequestSchema = z.object({
  ClusterId: z.number().positive(),
  QueueName: z.string().min(1),
  Create: z.boolean().default(false),
  Type: z.union([z.literal("classic"), z.literal("quorum")]),
});
