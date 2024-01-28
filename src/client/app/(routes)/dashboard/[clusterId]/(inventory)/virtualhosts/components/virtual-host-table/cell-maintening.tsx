"use client";

import { Button } from "@/components/ui/button";
import { RabbitMqExchange } from "@/models/exchange";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import { BadgeCheck, Hammer, RefreshCcw } from "lucide-react";
import React from "react";

interface CellMainteningProps {
  Data: RabbitMqVirtualHost;
}
function CellMaintening({ Data }: CellMainteningProps) {
  if (Data.IsInCluster && !Data.IsInDatabase) {
    return (
      <Button size="sm" variant="destructive">
        {<Hammer className="w-4 h-4 fill-white mx-1" />}
        {"Not Tracked"}
      </Button>
    );
  }

  if (Data.IsInDatabase && Data.IsInCluster) {
    return (
      <Button size="sm" variant="success">
        {<BadgeCheck className="w-4 h-4  mx-1" />}
        {"Tracked"}
      </Button>
    );
  }

  if (Data.IsInDatabase && !Data.IsInCluster) {
    return (
      <Button size="sm" variant="warn">
        {<RefreshCcw className="w-4 h-4  mx-1" />}
        <p className="text-muted-foreground">Out of Sync</p>
      </Button>
    );
  }

  return <p>ERROR</p>;
}

export default CellMaintening;
