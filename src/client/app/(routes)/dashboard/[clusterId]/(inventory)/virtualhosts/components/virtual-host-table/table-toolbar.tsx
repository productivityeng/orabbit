"use client";

import { Table } from "@tanstack/react-table";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { FileStack, RefreshCcwDot, XCircle } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import LockItem from "@/components/lock-item/lock-item";
import _ from "lodash";
import { GetActiveLocker } from "@/lib/utils";

import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import { useContext } from "react";
import { VirtualTableContext } from "./virtualhost-table-context";

interface DataTableToolbarProps {
  table: Table<RabbitMqVirtualHost>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {
  const {
    OnImportVirtualHostClick,
    OnRemoveTrackingFromVirtualHost,
    OnSyncronizeVirtualHost,
    HandleLockItem,
  } = useContext(VirtualTableContext);
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
  const router = useRouter();

  let selectedVirtualHost: RabbitMqVirtualHost | null = null;
  if (isRowSelected) {
    selectedVirtualHost = table.getFilteredSelectedRowModel().rows[0].original;
  }

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
        <Button
          size="sm"
          disabled={IsImporDisabled}
          data-testid="import-vhost-button"
          onClick={() =>
            selectedVirtualHost &&
            OnImportVirtualHostClick?.(selectedVirtualHost)
          }
        >
          <FileStack className="w-4 h-4 mr-2" /> Importar
        </Button>
        <Button
          onClick={() =>
            selectedVirtualHost &&
            OnSyncronizeVirtualHost?.(selectedVirtualHost)
          }
          size="sm"
          data-testid="sync-vhost-button"
          disabled={IsSyncronizeDisable}
          className="h-8"
        >
          <RefreshCcwDot className="w-4 h-4 mr-2" /> Sincronizar
        </Button>
        <LockItem
          Disabled={IsLockDisabled}
          onLockItem={({ reason }) =>
            selectedVirtualHost && HandleLockItem?.(selectedVirtualHost, reason)
          }
          Label={`fila ${selectedVirtualHost?.Name}`}
          Lockers={selectedVirtualHost?.Lockers}
        />

        <Button
          onClick={() =>
            selectedVirtualHost &&
            OnRemoveTrackingFromVirtualHost?.(selectedVirtualHost)
          }
          data-testid="remove-vhost-button"
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
