"use client";

import { Table } from "@tanstack/react-table";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  FileStack,
  RefreshCcwDot,
  XCircle,
} from "lucide-react";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import LockItem from "@/components/lock-item/lock-item";
import _ from "lodash";
import { CreateLockerAction } from "@/actions/locker";
import { z } from "zod";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";
import { GetActiveLocker } from "@/lib/utils";
import { RabbitMqUser } from "@/models/users";
import { useContext } from "react";
import { UserTableContext } from "./user-table-context";

interface DataTableToolbarProps {
  table: Table<RabbitMqUser>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {
  const router = useRouter();
  const {onSyncronizeUser,onRemoveUser,onImportUser} = useContext(UserTableContext);
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;

  let selectUser: RabbitMqUser | null = null;
  if (isRowSelected) {
    selectUser = table.getFilteredSelectedRowModel().rows[0].original;
  }




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
          data-testid="import-user-button"
          onClick={async () => {
            const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
            if (!isRowSelected) return;
            const selectUser = table.getFilteredSelectedRowModel().rows[0].original;
            await onImportUser?.(selectUser);
          }}
        >
          <FileStack className="w-4 h-4 mr-2" /> Importar
        </Button>
        <Button
          onClick={async () => {
            const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
            if (!isRowSelected) return;
            const selectUser = table.getFilteredSelectedRowModel().rows[0].original;
            await onSyncronizeUser?.(selectUser);
          }}
          size="sm"
          data-testid="syncronize-user-button"
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
          onClick={async () => {
            const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
            if (!isRowSelected) return;
            const selectUser = table.getFilteredSelectedRowModel().rows[0].original;
            await onRemoveUser?.(selectUser);
          
          }}
          size="sm"
          variant="destructive"
          disabled={IsRemoveDisable}
          data-testid="remove-user-button"
          className="h-8"
        >
          <XCircle className="w-4 h-4 mr-2" /> Remover
        </Button>
      </div>
    </div>
  );
}
