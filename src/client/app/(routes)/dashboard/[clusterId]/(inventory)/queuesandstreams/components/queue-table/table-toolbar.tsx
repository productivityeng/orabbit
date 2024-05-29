"use client";

import { Table } from "@tanstack/react-table";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Compass,
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
import Link from "next/link";
import { useContext } from "react";
import { QueueTableContext } from "./queue-table-context";

interface DataTableToolbarProps {
  table: Table<RabbitMqQueue>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {

  const {onSyncronizeQueueClick,ClusterId,onRemoveQueueClick,onImportQueueClick} = useContext(QueueTableContext)
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
  const router = useRouter();

  let selectedQueue: RabbitMqQueue | null = null;
  if (isRowSelected) {
    selectedQueue = table.getFilteredSelectedRowModel().rows[0].original;
  }
 

  const onLockItem = async (data: z.infer<typeof LockItemFormSchema>) => {
    if (!selectedQueue) return;
    let toastId = toast.loading(`Trancando fila ${selectedQueue.Name} ...`);

    try {
      await CreateLockerAction(
        selectedQueue.ClusterId,
        "queue",
        selectedQueue.ID,
        {
          reason: data.reason,
          responsible: "victor@riskskipper.com",
        }
      );
      toast.success(`Fila ${selectedQueue.Name} trancada`, {
        id: toastId,
      });
      router.refresh();
      table.toggleAllRowsSelected(false);
    } catch (error) {
      toast.error(
        `Erro ${JSON.stringify(error)} ao trancar fila ${selectedQueue.Name}`,
        {
          id: toastId,
        }
      );
    }
  };

  const IsImporDisabled = !isRowSelected || selectedQueue?.IsInDatabase;

  const IsRemoveDisable = !isRowSelected || !selectedQueue?.IsInDatabase;

  const IsSyncronizeDisable =
    !isRowSelected ||
    selectedQueue?.IsInCluster ||
    !selectedQueue?.IsInDatabase;

  const IsLockDisabled =
    !selectedQueue?.IsInDatabase ||
    GetActiveLocker(selectedQueue?.Lockers) != null;

  const IsManageDisable = !selectedQueue?.IsInDatabase;

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
        <Button
          size="sm"
          disabled={IsImporDisabled}
          onClick={() => selectedQueue && onImportQueueClick?.(selectedQueue)}
        >
          <FileStack className="w-4 h-4 mr-2" /> Importar
        </Button>
        <Button
          onClick={() => selectedQueue && onSyncronizeQueueClick?.(selectedQueue)}
          size="sm"
          disabled={IsSyncronizeDisable}
          className="h-8"
        >
          <RefreshCcwDot className="w-4 h-4 mr-2" /> Sincronizar
        </Button>
        <Link
          href={`/dashboard/${ClusterId}/queuesandstreams/${selectedQueue?.ID}`}
        >
          <Button
            disabled={IsManageDisable}
            size="sm"
            className="h-8 bg-sky-500 text-white"
          >
            <Compass className="w-4 h-4 mr-2" />
            Gerenciar
          </Button>
        </Link>

        <LockItem
          Disabled={IsLockDisabled}
          onLockItem={onLockItem}
          Label={`fila ${selectedQueue?.Name}`}
          Lockers={selectedQueue?.Lockers}
        />

        <Button
          onClick={() => selectedQueue && onRemoveQueueClick?.(selectedQueue)}
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
