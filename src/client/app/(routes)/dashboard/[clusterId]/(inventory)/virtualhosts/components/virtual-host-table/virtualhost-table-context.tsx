import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import React from "react";
import { z } from "zod";

interface VirtualTableContext {
  OnImportVirtualHostClick?: (vhost: RabbitMqVirtualHost) => Promise<void>;
  OnRemoveTrackingFromVirtualHost?: (
    vhost: RabbitMqVirtualHost
  ) => Promise<void>;
  OnSyncronizeVirtualHost?: (vhost: RabbitMqVirtualHost) => Promise<void>;
  HandleLockItem?: (
    virtualHost: RabbitMqVirtualHost,
    reason: string
  ) => Promise<void>;
}
export const VirtualTableContext = React.createContext<VirtualTableContext>({});
