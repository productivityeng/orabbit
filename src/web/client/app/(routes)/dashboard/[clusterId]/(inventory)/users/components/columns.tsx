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
    accessorKey: "type",
    header: "Type",
  },
  {
    accessorKey: "isRegistered",
    header: () => <b>Is Registered</b>,
    cell: ({ row }) => <CellMaintening User={row.original} />,
  },
  {
    id: "actions",
    header: () => <b>Actions</b>,
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
