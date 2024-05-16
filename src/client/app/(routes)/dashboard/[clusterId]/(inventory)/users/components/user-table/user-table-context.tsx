import { FrontResponse } from "@/actions/common/frontresponse";
import { RabbitMqUser } from "@/models/users";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";
import React from "react";
import { z } from "zod";

export interface UserTableContextProps {
    onSyncronizeUser?(user:RabbitMqUser): Promise<void>;
    onRemoveUser?(user:RabbitMqUser): Promise<void>;
    onImportUser?(user:RabbitMqUser): Promise<void>;
    onLockUser?(user: RabbitMqUser,data: z.infer<typeof LockItemFormSchema>): Promise<void>;
    onUnlockUser?(user: RabbitMqUser,lockId: number): Promise<void>;
}
export const UserTableContext = React.createContext<UserTableContextProps>({});
