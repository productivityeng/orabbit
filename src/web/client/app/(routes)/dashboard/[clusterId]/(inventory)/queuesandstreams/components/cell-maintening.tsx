"use client";

import { importUserFromCluster } from "@/actions/users";
import { Button } from "@/components/ui/button";
import { RabbitMqQueue } from "@/models/queues";
import { BadgeCheck, Frown, Hammer, RefreshCcw } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import React from "react";
import { toast } from "react-hot-toast";

interface CellMainteningProps {
  Queue: RabbitMqQueue;
}
function CellMaintening({ Queue }: CellMainteningProps) {
  if (Queue.IsInCluster && !Queue.IsInDatabase) {
    return (
      <Button size="sm" variant="destructive">
        {<Hammer className="w-4 h-4 fill-white mx-1" />}
        {"Not Tracked"}
      </Button>
    );
  }

  if (Queue.IsInDatabase && Queue.IsInCluster) {
    return (
      <Button size="sm" variant="success">
        {<BadgeCheck className="w-4 h-4  mx-1" />}
        {"Tracked"}
      </Button>
    );
  }

  if (Queue.IsInDatabase && !Queue.IsInCluster) {
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
