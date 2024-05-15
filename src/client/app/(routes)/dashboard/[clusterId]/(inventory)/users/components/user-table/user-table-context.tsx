import { FrontResponse } from "@/actions/common/frontresponse";
import { RabbitMqUser } from "@/models/users";
import React from "react";

export interface UserTableContextProps {
    onSyncronizeUser(user:RabbitMqUser): Promise<void>;
    onRemoveUser(user:RabbitMqUser): Promise<void>;
}
export const UserTableContext = React.createContext<UserTableContextProps>({    
    onSyncronizeUser: async (user:RabbitMqUser) => {
        throw new Error("SyncronizeUserAction not implemented");
    },
    onRemoveUser: async (user:RabbitMqUser) => {
        throw new Error("RemoveUserAction not implemented");
    }
});
