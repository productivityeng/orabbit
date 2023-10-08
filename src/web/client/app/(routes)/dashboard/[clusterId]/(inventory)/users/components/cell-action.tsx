import React, { useState } from "react";
import toast from "react-hot-toast";
import { useParams, useRouter } from "next/navigation";
import AlertModal from "@/components/Modals/alert-danger-modal";
import { Edit, Files, RefreshCw, SettingsIcon, Trash } from "lucide-react";
import { UserColumn } from "./columns";
import {
  SyncronizeUserAction,
  deleteUserFromRabbit,
  importUserFromCluster,
} from "@/actions/users";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";
import { RabbitMqUser } from "@/models/users";

interface CellActionProps {
  data: RabbitMqUser;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const removeUserFromCluster = async () => {
    const toastId = toast.loading(<p>Deletando ususario {data.Username}</p>);
    let result = await deleteUserFromRabbit(Number(params.clusterId), data.Id);

    if (result.ErrorMessage) {
      toast.error(<p>Erro ao deletar ususario {data.Username}</p>, {
        id: toastId,
      });
      return;
    }
    toast.success(<p>Usuario {data.Username} deletado com sucesso!</p>, {
      id: toastId,
    });
    router.refresh();
  };

  const importUser = async () => {
    const toastId = toast.loading(<p>Importando ususario {data.Username}</p>);
    let result = await importUserFromCluster({
      ClusterId: Number(params.clusterId),
      Create: false,
      Username: data.Username,
    });
    if (result.ErrorMessage) {
      toast.error(<p>Erro ao importar ususario {data.Username}</p>, {
        id: toastId,
      });
      return;
    }
    toast.success(<p>Usuario {data.Username} importado com sucesso!</p>, {
      id: toastId,
    });
    router.refresh();
  };

  const syncronizeUser = async () => {
    const toastId = toast.loading(
      <p>Sincronizando ususario {data.Username}</p>
    );
    let result = await SyncronizeUserAction(Number(params.clusterId), data.Id);
    if (result.ErrorMessage) {
      toast.error(<p>Erro ao sincronizar ususario {data.Username}</p>, {
        id: toastId,
      });
      return;
    }
    toast.success(<p>Usuario {data.Username} sincronizado com sucesso!</p>, {
      id: toastId,
    });
    router.refresh();
  };

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button
            variant={"ghost"}
            className="w-8 h-8 p-0 focus-visible:ring-0  focus-visible:ring-offset-0"
          >
            <SettingsIcon
              className={cn("w-4 h-4 duration-200 ease-in-out ", {
                "text-rabbit w-8 h-8": false,
              })}
            />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Actions</DropdownMenuLabel>

          {data.IsInCluster && (
            <DropdownMenuItem onClick={removeUserFromCluster}>
              <Edit className="mr-2 h-4 w-4" /> Remover do cluster
            </DropdownMenuItem>
          )}
          <DropdownMenuItem onClick={importUser}>
            <Files className="mr-2 h-4 w-4" /> Importar
          </DropdownMenuItem>
          {data.IsInDatabase && !data.IsInCluster && (
            <DropdownMenuItem onClick={syncronizeUser}>
              <RefreshCw className="mr-2 h-4 w-4" /> Sincronizar
            </DropdownMenuItem>
          )}
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
}

export default CellAction;
