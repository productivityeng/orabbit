import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import React from "react";

interface VirtualTableContext {
  OnImportVirtualHostClick?: (vhost: RabbitMqVirtualHost) => Promise<void>;
}
export const VirtualTableContext = React.createContext<VirtualTableContext>({});
