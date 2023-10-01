"use client";

import { deleteUserFromTracking, importUserFromCluster } from "@/actions/users";
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
      await deleteUserFromTracking(Number(params["clusterId"]), User.Id);
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

  return (
    <>
      {!User.IsRegistered ? (
        <Button
          onClick={handleImportUser}
          size="sm"
          variant="destructive"
          className="hover:bg-green-500"
        >
          {<Hammer className="w-4 h-4 fill-white mx-1" />}
          {"Import"}
        </Button>
      ) : (
        <Button
          onClick={handleDetachUser}
          size="sm"
          variant="success"
          className="hover:bg-red-500"
        >
          {<BadgeCheck className="w-4 h-4  mx-1" />}
          {"Tracked"}
        </Button>
      )}
    </>
  );
}

export default CellMaintening;
