"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellAction from "./cell-action";
import CellMaintening from "./cell-maintening";
import { RabbitMqUser } from "@/types";

export const UserColumn: ColumnDef<RabbitMqUser>[] = [
  {
    accessorKey: "Id",
    header: "Id",
  },
  {
    accessorKey: "Username",
    header: "Username",
  },

  {
    accessorKey: "isRegistered",
    header: () => <b>Is Registered</b>,
    cell: ({ row }) => <CellMaintening User={row.original} />,
  },
  {
    id: "actions",
    header: () => <b></b>,
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
