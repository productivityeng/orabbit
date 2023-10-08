import React, { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import {
  Copy,
  Edit,
  Files,
  MoreHorizontal,
  RefreshCcw,
  RefreshCw,
  SettingsIcon,
  Trash,
} from "lucide-react";
import toast from "react-hot-toast";
import { useParams, useRouter } from "next/navigation";
import { RabbitMqQueue } from "@/models/queues";
import { cn } from "@/lib/utils";
import {
  removeQueueFromClusterAction,
  syncronizeQueueAction,
} from "@/actions/queue";

interface CellActionProps {
  data: RabbitMqQueue;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const syncronizeQueue = async () => {
    const toastId = toast.loading(
      <p>
        Sincronizando fila <p className="text-rabbit">{data.Name}</p> ...
      </p>
    );
    try {
      let response = await syncronizeQueueAction({
        ClusterId: Number(params.clusterId),
        QueueId: data.ID,
      });
      if (response.ErrorMessage) {
        toast.error(
          <p>
            Erro {response.ErrorMessage} ao sincronizar fila{" "}
            <p className="text-rabbit">{data.Name}</p>{" "}
          </p>,
          {
            id: toastId,
          }
        );
        return;
      }

      toast.success(
        <p>
          Fila <p className="text-rabbit">{data.Name}</p> sincronizada com
          sucesso{" "}
        </p>,
        {
          id: toastId,
        }
      );
      router.refresh();
    } catch (error) {
      toast.error(
        <p>
          Erro ao sincronizar fila <p className="text-rabbit">{data.Name}</p>{" "}
        </p>,
        {
          id: toastId,
        }
      );
    }
  };

  const removeQueueFromCluster = async () => {
    const toastId = toast.loading(
      <p>
        Removendo fila <p className="text-rabbit">{data.Name}</p> ...
      </p>
    );
    try {
      let response = await removeQueueFromClusterAction({
        ClusterId: Number(params.clusterId),
        QueueId: data.ID,
      });
      if (response.ErrorMessage) {
        toast.error(
          <p>
            Erro {response.ErrorMessage} ao remover fila{" "}
            <p className="text-rabbit">{data.Name}</p>{" "}
          </p>,
          {
            id: toastId,
          }
        );
        return;
      }

      toast.success(
        <p>
          Fila <p className="text-rabbit">{data.Name}</p> removida com sucesso
        </p>,
        {
          id: toastId,
        }
      );
      router.refresh();
    } catch (error) {
      toast.error(
        <p>
          Erro ao remover fila <p className="text-rabbit">{data.Name}</p>{" "}
        </p>,
        {
          id: toastId,
        }
      );
    }
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
            <DropdownMenuItem>
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
