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
import toast from "react-hot-toast";
import { some } from "lodash";
import LockItem from "@/components/lock-item/lock-item";
import _ from "lodash";
import { CreateLockerAction } from "@/actions/locker";
import { z } from "zod";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";
import { GetActiveLocker } from "@/lib/utils";

import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import {
  removeVirtualHostAction,
  syncronizeVirtualHostAction,
} from "@/actions/virtualhost";

interface DataTableToolbarProps {
  table: Table<RabbitMqVirtualHost>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {
  const { clusterId } = useParams() as { clusterId: string };
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
  const router = useRouter();

  let selectedVirtualHost: RabbitMqVirtualHost | null = null;
  if (isRowSelected) {
    selectedVirtualHost = table.getFilteredSelectedRowModel().rows[0].original;
  }

  const onSyncronizeClick = async () => {
    if (!selectedVirtualHost) return;
    let toastId = toast.loading("Sincronizando VirtualHost...");
    try {
      let result = await syncronizeVirtualHostAction(
        Number(clusterId),
        selectedVirtualHost.Id
      );
      if (result.Result) {
        toast.success("VirtualHost sincronizado com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error(
          `Falha ao sincronizar VirtualHost: ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(`Falha ao sincronizar VirtualHost:${error}`, {
        id: toastId,
      });
    }
  };

  const onImportClick = async () => {};

  const onRemoveClick = async () => {
    if (!selectedVirtualHost) return;
    let toastId = toast.loading("Removendo VirtualHost...");
    try {
      let result = await removeVirtualHostAction(
        Number(clusterId),
        selectedVirtualHost.Id
      );
      if (result.Result) {
        toast.success("VirtualHost removido com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error(`Falha ao remover VirtualHost: ${result.ErrorMessage}`, {
          id: toastId,
        });
      }
    } catch (error) {
      toast.error(`Falha ao remover VirtualHost:${error}`, { id: toastId });
    }
  };

  const onLockItem = async (data: z.infer<typeof LockItemFormSchema>) => {
    if (!selectedVirtualHost) return;
    let toastId = toast.loading("Travando VirtualHost...");
    try {
      let result = await CreateLockerAction(
        Number(clusterId),
        "virtualhost",
        selectedVirtualHost.Id,
        {
          reason: data.reason,
          responsible: "Victor",
        }
      );
      if (result.Result) {
        toast.success("VirtualHost bloqueado com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error("Falha ao travar VirtualHost", { id: toastId });
      }
    } catch (error) {
      toast.error("Falha ao travar VirtualHost", { id: toastId });
    }
  };

  const IsImporDisabled = !isRowSelected || selectedVirtualHost?.IsInDatabase;

  const IsRemoveDisable =
    !isRowSelected ||
    !selectedVirtualHost?.IsInCluster ||
    !selectedVirtualHost?.IsInDatabase;

  const IsSyncronizeDisable =
    !isRowSelected ||
    selectedVirtualHost?.IsInCluster ||
    !selectedVirtualHost?.IsInDatabase;

  const IsLockDisabled =
    !selectedVirtualHost?.IsInDatabase ||
    GetActiveLocker(selectedVirtualHost?.Lockers) != null;

  return (
    <div className="flex items-center justify-between">
      <div className="flex flex-1 items-center space-x-2">
        <Input
          placeholder="Filtrar fila"
          onChange={(event) => {
            table.setGlobalFilter(event.target.value);
          }}
          className="h-8 w-[150px] md:w-[250px]"
        />
        <Button size="sm" disabled={IsImporDisabled} onClick={onImportClick}>
          <FileStack className="w-4 h-4 mr-2" /> Importar
        </Button>
        <Button
          onClick={onSyncronizeClick}
          size="sm"
          disabled={IsSyncronizeDisable}
          className="h-8"
        >
          <RefreshCcwDot className="w-4 h-4 mr-2" /> Sincronizar
        </Button>
        <LockItem
          Disabled={IsLockDisabled}
          onLockItem={onLockItem}
          Label={`fila ${selectedVirtualHost?.Name}`}
          Lockers={selectedVirtualHost?.Lockers}
        />

        <Button
          onClick={onRemoveClick}
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
