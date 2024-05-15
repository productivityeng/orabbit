import { FrontResponse } from "@/actions/common/frontresponse";
import { RabbitMqUser } from "@/models/users";
import React from "react";

export interface UserTableContextProps {
    onSyncronizeUserClick(user:RabbitMqUser): Promise<void>;
}
export const UserTableContext = React.createContext<UserTableContextProps>({    
    onSyncronizeUserClick: async (user:RabbitMqUser) => {
        throw new Error("SyncronizeUserAction not implemented");
    }
});
