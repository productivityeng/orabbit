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
import { useTranslations } from "next-intl";

interface DataTableToolbarProps {
  table: Table<RabbitMqUser>;
}

export function DataTableToolbar({ table }: DataTableToolbarProps) {
  const {onSyncronizeUser,onRemoveUser,onImportUser,onLockUser} = useContext(UserTableContext);
  const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
  const t = useTranslations("Dashboard.UsersPage");
  
  let selectUser: RabbitMqUser | null = null;
  if (isRowSelected) {
    selectUser = table.getFilteredSelectedRowModel().rows[0].original;
  }

  const hasActiveLocker = selectUser && GetActiveLocker(selectUser.Lockers) != null

  const IsImporDisabled = hasActiveLocker || !isRowSelected || selectUser?.IsInDatabase;

  const IsRemoveDisabled = hasActiveLocker ||!isRowSelected || !selectUser?.IsInCluster;

  const IsSyncronizeDisable = hasActiveLocker || !isRowSelected || selectUser?.IsInCluster || !selectUser?.IsInDatabase;

  const IsLockDisabled = hasActiveLocker || !selectUser?.IsInDatabase 
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
          <FileStack className="w-4 h-4 mr-2" /> {t("Import")}
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
          <RefreshCcwDot className="w-4 h-4 mr-2" /> {t("Syncronize")}
        </Button>
        <LockItem
          Disabled={IsLockDisabled}
          onLockItem={async (data: z.infer<typeof LockItemFormSchema>) => {
            const isRowSelected = table.getFilteredSelectedRowModel().rows.length > 0;
            if (!isRowSelected) return;
            const selectUser = table.getFilteredSelectedRowModel().rows[0].original;
            await onLockUser?.(selectUser,data );
          }}
          Label={`${t("Queue")} ${selectUser?.Username}`}
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
          disabled={IsRemoveDisabled}
          data-testid="remove-user-button"
          className="h-8"
        >
          <XCircle className="w-4 h-4 mr-2" /> {t("Remove")}
        </Button>
      </div>
    </div>
  );
}
