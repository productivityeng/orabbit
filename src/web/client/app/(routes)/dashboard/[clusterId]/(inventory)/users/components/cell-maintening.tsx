"use client";

import { deleteUserFromCluster, importUserFromCluster } from "@/actions/users";
import { Button } from "@/components/ui/button";
import { RabbitMqUser } from "@/models/users";
import { BadgeCheck, Frown, Hammer, RefreshCcw } from "lucide-react";
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
      toast.success(
        <p>
          Usuario {<b className="text-rabbit">{User.Username}</b>} importado com
          sucesso
        </p>,
        {
          id: toastId,
        }
      );
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

  const handleDetachUser = async () => {
    const toastId = toast.loading("Removendo usuario do ostern...");
    try {
      await deleteUserFromCluster(Number(params["clusterId"]), User.Id);
      toast.success(
        <p>
          Usuario {<b className="text-rabbit">{User.Username}</b>} removido do
          ostern com sucesso
        </p>,
        {
          id: toastId,
        }
      );
      route.refresh();
    } catch (error) {
      toast.error(
        <p>
          Something wen't wrong <Frown />
        </p>,
        {
          id: toastId,
        }
      );
    }
  };

  if (User.IsInCluster && User.IsInDatabase) {
    return (
      <Button size="sm" variant="success">
        {<BadgeCheck className="w-4 h-4  mx-1" />}
        {"Tracked"}
      </Button>
    );
  }

  if (User.IsInDatabase && !User.IsInCluster) {
    return (
      <Button size="sm" variant="warn">
        {<RefreshCcw className="w-4 h-4  mx-1" />}
        <p className="text-muted-foreground">Out of Sync</p>
      </Button>
    );
  }

  if (User.IsInCluster && !User.IsInDatabase) {
    return (
      <Button size="sm" variant="destructive">
        {<Hammer className="w-4 h-4 fill-white mx-1" />}
        {"Not Tracked"}
      </Button>
    );
  }
}

export default CellMaintening;
