import React, { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { Edit, Files, RefreshCw, SettingsIcon } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import { cn, standardToastableAction } from "@/lib/utils";
import {
  ImportQueueFromClusterAction,
  removeQueueFromClusterAction,
} from "@/actions/queue";
import { VirtualHosts } from "../../models/virtualhosts";

interface VirtualHostCellActionProps {
  data: VirtualHosts;
}

function VirtualHostCellAction({ data }: VirtualHostCellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const syncronizeQueue = async () => {
    // await standardToastableAction(
    //   async () => {
    //     await syncronizeQueueAction({
    //       ClusterId: Number(params.clusterId),
    //       QueueId: data.Id,
    //     });
    //   },
    //   <p>
    //     Sincronizando fila <span className="text-rabbit">{data.Name}</span> ...
    //   </p>,
    //   <p>
    //     Fila <span className="text-rabbit">{data.Name}</span> sincronizada com
    //     sucesso{" "}
    //   </p>,
    //   <p>
    //     Erro ao sincronizar fila{" "}
    //     <span className="text-rabbit">{data.Name}</span>{" "}
    //   </p>,
    //   [
    //     () => {
    //       router.refresh();
    //     },
    //   ],
    //   []
    // );
  };

  const removeQueueFromCluster = async () => {
    // await standardToastableAction(
    //   async () => {
    //     await removeQueueFromClusterAction({
    //       ClusterId: Number(params.clusterId),
    //       QueueId: data.ID,
    //     });
    //   },
    //   <p>
    //     Removendo fila <span className="text-rabbit">{data.Name}</span> ...
    //   </p>,
    //   <p>
    //     Fila <span className="text-rabbit">{data.Name}</span> removida com
    //     sucesso
    //   </p>,
    //   <p>
    //     Erro ao remover fila <span className="text-rabbit">{data.Name}</span>{" "}
    //   </p>,
    //   [
    //     () => {
    //       router.refresh();
    //     },
    //   ],
    //   []
    // );
  };

  const importQueueFromCluster = async () => {
    await standardToastableAction(
      async () => {
        await ImportQueueFromClusterAction(Number(params.clusterId), data.Name);
      },
      <p>Importando virtualHost do cluster</p>,
      <p>VirtualHost importada com sucesso!</p>,
      <p>VirtualHost importada com sucesso!</p>,
      [
        () => {
          router.refresh();
        },
      ],
      []
    );
  };

  return (
    <DropdownMenu onOpenChange={setIsMenuOpen}>
      <DropdownMenuTrigger asChild>
        <Button variant={"ghost"} className="w-8 h-8 p-0 active:scale-100">
          <SettingsIcon
            className={cn("w-6 h-6 duration-200 ease-in-out ", {
              "text-rabbit ": isMenuOpen,
            })}
          />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel>Actions</DropdownMenuLabel>
        {data.IsInDatabase && data.IsInCluster && (
          <DropdownMenuItem onClick={removeQueueFromCluster}>
            <Edit className="mr-2 h-4 w-4" /> Remove From Cluster
          </DropdownMenuItem>
        )}
        {data.IsInCluster && !data.IsInDatabase && (
          <DropdownMenuItem onClick={importQueueFromCluster}>
            <Files className="mr-2 h-4 w-4" /> Importar
          </DropdownMenuItem>
        )}
        {data.IsInDatabase && !data.IsInCluster && (
          <DropdownMenuItem onClick={syncronizeQueue}>
            <RefreshCw className="mr-2 h-4 w-4" /> Syncronize
          </DropdownMenuItem>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export default VirtualHostCellAction;
