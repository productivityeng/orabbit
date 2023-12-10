"use client";

import { Button } from "@/components/ui/button";
import { BadgeCheck, Hammer, RefreshCcw } from "lucide-react";
import React from "react";
import { VirtualHosts } from "../../models/virtualhosts";

interface CellVirtualHostMainteningProps {
  VirtualHost: VirtualHosts;
}
function CellVirtualHostMaintening({
  VirtualHost,
}: CellVirtualHostMainteningProps) {
  if (VirtualHost.IsInCluster && !VirtualHost.IsInDatabase) {
    return (
      <Button size="sm" variant="destructive">
        {<Hammer className="w-4 h-4 fill-white mx-1" />}
        {"Not Tracked"}
      </Button>
    );
  }

  if (VirtualHost.IsInDatabase && VirtualHost.IsInCluster) {
    return (
      <Button size="sm" variant="success">
        {<BadgeCheck className="w-4 h-4  mx-1" />}
        {"Tracked"}
      </Button>
    );
  }

  if (VirtualHost.IsInDatabase && !VirtualHost.IsInCluster) {
    return (
      <Button size="sm" variant="warn">
        {<RefreshCcw className="w-4 h-4  mx-1" />}
        <p className="text-muted-foreground">Out of Sync</p>
      </Button>
    );
  }

  return <p>ERROR</p>;
}

export default CellVirtualHostMaintening;
