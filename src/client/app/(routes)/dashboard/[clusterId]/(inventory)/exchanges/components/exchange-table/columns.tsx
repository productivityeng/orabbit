"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellAction from "./cell-action";
import CellMaintening from "./cell-maintening";
import CellArguments from "./cell-arguments";
import { RabbitMqQueue } from "@/models/queues";
import CellDurable from "./cell-durable";
import { Checkbox } from "@/components/ui/checkbox";
import { GetActiveLocker } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { LockIcon, UnlockIcon } from "lucide-react";
import CellLocker from "./cell-locker";
import { RabbitMqExchange } from "@/models/exchange";

export const RabbitMqExchangeColumnDef: ColumnDef<RabbitMqExchange>[] = [
  {
    id: "select",
    header: ({ table }) => (
      <Checkbox
        checked={
          table.getIsAllPageRowsSelected() ||
          (table.getIsSomePageRowsSelected() && "indeterminate")
        }
        onCheckedChange={(value) => table.toggleAllPageRowsSelected(!!value)}
        aria-label="Select all"
        className="translate-y-[2px]"
      />
    ),
    cell: ({ row }) => (
      <Checkbox
        checked={row.getIsSelected()}
        onCheckedChange={(value) => row.toggleSelected(!!value)}
        aria-label="Select row"
        className="translate-y-[2px]"
      />
    ),
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "Id",
    header: "ID",
  },
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "Type",
    header: "Type",
  },
  {
    accessorKey: "Lockers",
    header: () => <b>Locked</b>,
    cell: ({ row }) => <CellLocker data={row.original} />,
  },
  {
    accessorKey: "IsRegistered",
    header: () => <b>Is Registered</b>,
    cell: ({ row }) => <CellMaintening Data={row.original} />,
  },
  {
    accessorKey: "Durable",
    header: () => <b>Durable</b>,
    cell: ({ row }) => <CellDurable Data={row.original} />,
  },
  {
    accessorKey: "Arguments",
    header: () => <b>Arguments</b>,
    cell: ({ row }) => <CellArguments Queue={row.original} />,
  },
];
