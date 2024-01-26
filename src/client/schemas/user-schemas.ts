import { z } from "zod";


export const CreateRabbitmqUserSchema = z.object({
    username: z.string().min(1),
    clusterId: z.number().int().positive(),
    password: z.string().optional(),
    create: z.boolean()
}).refine((data) => data.create == true ? data.password != undefined: true,{
    message: 'Password is required when create is true',
    path: ["password"]
});

