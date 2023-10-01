"use client";

import { importUserFromCluster } from "@/actions/users";
import { Button } from "@/components/ui/button";
import { RabbitMqQueue, RabbitMqUser } from "@/types";
import { BadgeCheck, Frown, Hammer } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import React from "react";
import { toast } from "react-hot-toast";

interface CellMainteningProps {
  Queue: RabbitMqQueue;
}
function CellMaintening({ Queue }: CellMainteningProps) {
  const params = useParams();
  const route = useRouter();

  const handleImportUser = async () => {};

  return (
    <>
      {!Queue.IsRegistered ? (
        <Button onClick={handleImportUser} size="sm" variant="destructive">
          {<Hammer className="w-4 h-4 fill-white mx-1" />}
          {"Import"}
        </Button>
      ) : (
        <Button size="sm" variant="success">
          {<BadgeCheck className="w-4 h-4  mx-1" />}
          {"Tracked"}
        </Button>
      )}
    </>
  );
}

export default CellMaintening;
