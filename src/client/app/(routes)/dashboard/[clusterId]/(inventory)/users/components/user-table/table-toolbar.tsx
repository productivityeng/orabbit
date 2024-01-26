"use client";

import { Table } from "@tanstack/react-table";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Delete,
  FileStack,
  LockIcon,
  RefreshCcwDot,
  RemoveFormatting,
  XCircle,
} from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import { RabbitMqQueue } from "@/models/queues";
import {
  ImportQueueFromClusterAction,
  removeQueueFromClusterAction,
  syncronizeQueueAction,
} from "@/actions/queue";
import toast from "react-hot-toast";
import { some } from "lodash";
import LockItem from "@/components/lock-item/lock-item";
import _ from "lodash";
import { CreateLockerAction } from "@/actions/locker";
import { z } from "zod";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";
import { GetActiveLocker } from "@/lib/utils";
import { RabbitMqUser } from "@/models/users";
import {
  SyncronizeUserAction,
  importUserFromCluster,
  removeUserFromCluster,
} from "@/actions/users";

interface DataTableToolbarProps {
  table: Table<RabbitMqUser>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {
  const { clusterId } = useParams() as { clusterId: string };
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
  const router = useRouter();

  let selectUser: RabbitMqUser | null = null;
  if (isRowSelected) {
    selectUser = table.getFilteredSelectedRowModel().rows[0].original;
  }

  const onSyncronizeUserClick = async () => {
    if (!selectUser) return;
    const toastId = toast.loading(
      `Sincronizando usuario ${selectUser.Username}`
    );

    try {
      let result = await SyncronizeUserAction(
        selectUser.ClusterId,
        selectUser.Id
      );
      if (result.Result) {
        toast.success(
          `Usuario ${selectUser.Username} sincronizado com sucesso`,
          {
            id: toastId,
          }
        );
        router.refresh();
      } else {
        toast.error(
          `Error ao sincronizar usuario ${selectUser.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao sincronizar usuario ${selectUser.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  const onImportUserClick = async () => {
    if (!selectUser) return;
    const toastId = toast.loading(`Importando usuario ${selectUser.Username}`);

    try {
      let result = await importUserFromCluster({
        ClusterId: selectUser.ClusterId,
        Username: selectUser.Username,
        Create: false,
      });
      if (result.Result) {
        toast.success(`Usuario ${selectUser.Username} importado com sucesso`, {
          id: toastId,
        });
        table.toggleAllRowsSelected(false);
        router.refresh();
      } else {
        toast.error(
          `Error ao importar usuario ${selectUser.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao importar usuario ${selectUser.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  const onRemoveUserClick = async () => {
    if (!selectUser) return;
    const toastId = toast.loading(`Removendo usuario ${selectUser.Username}`);

    try {
      let result = await removeUserFromCluster(
        selectUser.ClusterId,
        selectUser.Id
      );
      if (result.Result) {
        toast.success(`Usuario ${selectUser.Username} removido com sucesso`, {
          id: toastId,
        });
        router.refresh();
      } else {
        toast.error(
          `Error ao remover usuario ${selectUser.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao remover usuario ${selectUser.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  const onLockItem = async (data: z.infer<typeof LockItemFormSchema>) => {
    if (!selectUser) return;
    const toastId = toast.loading(`Bloqueando usuario ${selectUser.Username}`);

    try {
      let result = await CreateLockerAction(
        selectUser.ClusterId,
        "user",
        selectUser.Id,
        {
          reason: data.reason,
          responsible: "Victor",
        }
      );
      if (result.Result) {
        toast.success(`Usuario ${selectUser.Username} bloqueado com sucesso`, {
          id: toastId,
        });
        router.refresh();
      } else {
        toast.error(
          `Error ao bloquear usuario ${selectUser.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao bloquear usuario ${selectUser.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  const IsImporDisabled = !isRowSelected || selectUser?.IsInDatabase;

  const IsRemoveDisable = !isRowSelected || !selectUser?.IsInCluster;

  const IsSyncronizeDisable =
    !isRowSelected || selectUser?.IsInCluster || !selectUser?.IsInDatabase;

  const IsLockDisabled =
    !selectUser?.IsInDatabase || GetActiveLocker(selectUser?.Lockers) != null;

  return (
    <div className="flex items-center justify-between">
      <div className="flex flex-1 items-center space-x-2">
        <Input
          placeholder="Filtrar usuario"
          onChange={(event) => {
            table.setGlobalFilter(event.target.value);
          }}
          className="h-8 w-[150px] md:w-[250px]"
        />
        <Button
          size="sm"
          disabled={IsImporDisabled}
          onClick={onImportUserClick}
        >
          <FileStack className="w-4 h-4 mr-2" /> Importar
        </Button>
        <Button
          onClick={onSyncronizeUserClick}
          size="sm"
          disabled={IsSyncronizeDisable}
          className="h-8"
        >
          <RefreshCcwDot className="w-4 h-4 mr-2" /> Sincronizar
        </Button>
        <LockItem
          Disabled={IsLockDisabled}
          onLockItem={onLockItem}
          Label={`fila ${selectUser?.Username}`}
          Lockers={selectUser?.Lockers}
        />

        <Button
          onClick={onRemoveUserClick}
          size="sm"
          variant="destructive"
          disabled={IsRemoveDisable}
          className="h-8"
        >
          <XCircle className="w-4 h-4 mr-2" /> Remover
        </Button>
      </div>
    </div>
  );
}
