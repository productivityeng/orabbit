"use client";

import { ColumnDef } from "@tanstack/react-table";
import { Checkbox } from "@/components/ui/checkbox";
import { RabbitMqUser } from "@/models/users";
import CellLocker from "./cell-locker";
import CellMaintening from "./cell-maintening";

export const RabbitMqUserTableColumnsDef: ColumnDef<RabbitMqUser>[] = [
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
        data-testid={`user-table-checkbox-${row.original.Id}`}
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
    accessorKey: "Username",
    header: "Username",
  },
  {
    header: "Is Registered",
    cell: ({ row }) => <CellMaintening User={row.original} />,
  },
  {
    accessorKey: "Lockers",
    header: () => <b>Locked</b>,
    cell: ({ row }) => <CellLocker User={row.original} />,
  },
  {
    accessorKey: "PasswordHash",
    header: "PasswordHash",
  },
];
