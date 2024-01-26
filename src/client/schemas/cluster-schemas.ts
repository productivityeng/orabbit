import { z } from "zod";


export const CreateRabbitMqClusterRequestSchema = z.object({
    description: z.string().min(1),
    host: z.string().min(1),
    name: z.string().min(1),
    password: z.string().min(5),
    port: z.number().int().positive(),
    user: z.string().min(1),
  });

