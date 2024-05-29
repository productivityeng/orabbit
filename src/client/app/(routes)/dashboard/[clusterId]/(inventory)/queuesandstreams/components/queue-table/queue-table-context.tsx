import { RabbitMqQueue } from "@/models/queues";
import React from "react";



interface QueueTableContextProps {
    RemoveQueueHandle? : () =>{}
    onSyncronizeQueueClick?(queue: RabbitMqQueue): Promise<void>
    onRemoveQueueClick?(queue:RabbitMqQueue): Promise<void>
    onImportQueueClick?(queue:RabbitMqQueue): Promise<void>
    ClusterId?: number
}

export const QueueTableContext =  React.createContext<QueueTableContextProps>({})