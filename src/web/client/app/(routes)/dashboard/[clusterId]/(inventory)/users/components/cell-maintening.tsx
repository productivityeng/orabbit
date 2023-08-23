"use client";

import { importUserFromCluster } from "@/actions/users";
import { Button } from "@/components/ui/button";
import { RabbitMqUser } from "@/types";
import { BadgeCheck, Frown, Hammer } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import React from "react";
import { toast } from "react-hot-toast";

interface CellMainteningProps {
  User: RabbitMqUser;
}
function CellMaintening({ User }: CellMainteningProps) {
  const params = useParams();
  const route = useRouter();

  const handleImportUser = async () => {
    const toastId = toast.loading("Importing user...");
    try {
      importUserFromCluster({
        ClusterId: Number(params["clusterId"]),
        Username: User.Username,
        Create: false,
      });
      toast.success("User imported!", {
        id: toastId,
      });
      route.refresh();
    } catch (error) {
      toast.error(
        <p>
          "Something wen't wrong " <Frown />
        </p>,
        {
          id: toastId,
        }
      );
    }
  };

  return (
    <>
      {!User.IsRegistered ? (
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
