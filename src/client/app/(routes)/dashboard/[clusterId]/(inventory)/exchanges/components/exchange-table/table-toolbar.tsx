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
import { RabbitMqExchange } from "@/models/exchange";
import {
  importExchangeFromClusterAction,
  removeExchangeFromClusterAction,
  syncronizeExchangeAction,
} from "@/actions/exchanges";

interface DataTableToolbarProps {
  table: Table<RabbitMqExchange>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {
  const { clusterId } = useParams() as { clusterId: string };
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
  const router = useRouter();

  let selectedExchange: RabbitMqExchange | null = null;
  if (isRowSelected) {
    selectedExchange = table.getFilteredSelectedRowModel().rows[0].original;
  }

  const onSyncronizeQueueClick = async () => {
    if (!selectedExchange) return;
    let toastId = toast.loading("Sincronizando fila...");
    try {
      let result = await syncronizeExchangeAction(
        Number(clusterId),
        selectedExchange.Id
      );
      if (result.Result) {
        toast.success("Exchange sincronizada com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error(
          `Falha ao sincronizar exchange com erro => ${result.ErrorMessage}`,
          { id: toastId }
        );
      }
    } catch (error) {
      toast.error("Falha ao sincronizar exchange", { id: toastId });
    }
  };

  const onImportClick = async () => {
    if (!selectedExchange) return;
    let toastId = toast.loading("Importando exchange...");
    try {
      let result = await importExchangeFromClusterAction(
        Number(clusterId),
        selectedExchange.Name,
        selectedExchange.VHost
      );
      if (result.Result) {
        toast.success("Exchange importada com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error(`Falha ao importar exchange: ${result.ErrorMessage}`, {
          id: toastId,
        });
      }
    } catch (error) {
      toast.error("Falha ao importar exchange", { id: toastId });
    }
  };

  const onRemoveClick = async () => {
    if (!selectedExchange) return;
    let toastId = toast.loading("Removendo exchange...");
    try {
      let result = await removeExchangeFromClusterAction(
        Number(clusterId),
        selectedExchange.Id
      );
      if (result.Result) {
        toast.success("Exchange removida com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error(`Falha ao remover exchange ${result.ErrorMessage}`, {
          id: toastId,
        });
      }
    } catch (error) {
      toast.error("Falha ao remover exchange", { id: toastId });
    }
  };

  const onLockItem = async (data: z.infer<typeof LockItemFormSchema>) => {
    if (!selectedExchange) return;
    let toastId = toast.loading("Travando Exchange...");
    try {
      let result = await CreateLockerAction(
        Number(clusterId),
        "exchange",
        selectedExchange.Id,
        {
          reason: data.reason,
          responsible: "Victor",
        }
      );
      if (result.Result) {
        toast.success("Exchange bloqueada com sucesso", { id: toastId });
        router.refresh();
      } else {
        toast.error("Falha ao travar Exchange", { id: toastId });
      }
    } catch (error) {
      toast.error("Falha ao travar Exchange", { id: toastId });
    }
  };

  const IsImporDisabled = !isRowSelected || selectedExchange?.IsInDatabase;

  const IsRemoveDisable =
    !isRowSelected ||
    !selectedExchange?.IsInCluster ||
    !selectedExchange?.IsInDatabase;

  const IsSyncronizeDisable =
    !isRowSelected ||
    selectedExchange?.IsInCluster ||
    !selectedExchange?.IsInDatabase;

  const IsLockDisabled =
    !selectedExchange?.IsInDatabase ||
    GetActiveLocker(selectedExchange?.Lockers) != null;

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
          onClick={onSyncronizeQueueClick}
          size="sm"
          disabled={IsSyncronizeDisable}
          className="h-8"
        >
          <RefreshCcwDot className="w-4 h-4 mr-2" /> Sincronizar
        </Button>
        <LockItem
          Disabled={IsLockDisabled}
          onLockItem={onLockItem}
          Label={`fila ${selectedExchange?.Name}`}
          Lockers={selectedExchange?.Lockers}
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
