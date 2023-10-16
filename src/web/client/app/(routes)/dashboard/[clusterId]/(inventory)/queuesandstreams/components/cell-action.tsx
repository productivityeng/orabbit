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
import { RabbitMqQueue } from "@/models/queues";
import { cn, standardToastableAction } from "@/lib/utils";
import {
  ImportQueueFromClusterAction,
  removeQueueFromClusterAction,
} from "@/actions/queue";

interface CellActionProps {
  data: RabbitMqQueue;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const syncronizeQueue = async () => {
    await standardToastableAction(
      async () => {
        await ImportQueueFromClusterAction(Number(params.clusterId), data.Name);
      },
      <p>
        Sincronizando fila <p className="text-rabbit">{data.Name}</p> ...
      </p>,
      <p>
        Fila <p className="text-rabbit">{data.Name}</p> sincronizada com sucesso{" "}
      </p>,
      <p>
        Erro ao sincronizar fila <p className="text-rabbit">{data.Name}</p>{" "}
      </p>,
      [
        () => {
          router.refresh();
        },
      ],
      []
    );
  };

  const removeQueueFromCluster = async () => {
    await standardToastableAction(
      async () => {
        await removeQueueFromClusterAction({
          ClusterId: Number(params.clusterId),
          QueueId: data.ID,
        });
      },
      <p>
        Removendo fila <p className="text-rabbit">{data.Name}</p> ...
      </p>,
      <p>
        Fila <p className="text-rabbit">{data.Name}</p> removida com sucesso
      </p>,
      <p>
        Erro ao remover fila <p className="text-rabbit">{data.Name}</p>{" "}
      </p>,
      [
        () => {
          router.refresh();
        },
      ],
      []
    );
  };

  const importQueueFromCluster = async () => {
    await standardToastableAction(
      async () => {
        await ImportQueueFromClusterAction(Number(params.clusterId), data.Name);
      },
      <p>"Importando fila do cluster"</p>,
      <p>"Fila importada com sucesso!"</p>,
      <p>"Fila importada com sucesso!"</p>,
      [
        () => {
          router.refresh();
        },
      ],
      []
    );
  };

  return (
    <>
      <DropdownMenu onOpenChange={setIsMenuOpen}>
        <DropdownMenuTrigger asChild>
          <Button
            variant={"ghost"}
            className="w-8 h-8 p-0 focus-visible:ring-0  focus-visible:ring-offset-0"
          >
            <SettingsIcon
              className={cn("w-4 h-4 duration-200 ease-in-out ", {
                "text-rabbit w-8 h-8": isMenuOpen,
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
    </>
  );
}

export default CellAction;
