"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellAction from "./cell-action";
import CellMaintening from "./cell-maintening";
import { RabbitMqUser } from "@/types";

export const columns: ColumnDef<RabbitMqUser>[] = [
  {
    accessorKey: "Username",
    header: "Username",
  },
  {
    accessorKey: "PasswordHash",
    header: "PasswordHash",
  },
  {
    accessorKey: "isRegistered",
    header: () => <b>Is Registered</b>,
    cell: ({ row }) => <CellMaintening User={row.original} />,
  },
  {
    id: "actions",
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
