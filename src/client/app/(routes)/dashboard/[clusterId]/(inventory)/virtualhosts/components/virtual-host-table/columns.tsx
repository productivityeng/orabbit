"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellMaintening from "./cell-maintening";
import DefaultQueueType from "./cell-default-queue-type";
import { Checkbox } from "@/components/ui/checkbox";
import CellLocker from "./cell-locker";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";

export const RabbitMqVirtualHostColumnDef: ColumnDef<RabbitMqVirtualHost>[] = [
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
        data-testid={`virtual-host-table-checkbox-${row.original.Id}`}
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
    accessorKey: "Arguments",
    header: () => <b>Default Queue Type</b>,
    cell: ({ row }) => <DefaultQueueType Data={row.original} />,
  },
];
