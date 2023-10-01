"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellAction from "./cell-action";
import CellMaintening from "./cell-maintening";
import { RabbitMqQueue } from "@/types";

export const RabbitMqQueueColumn: ColumnDef<RabbitMqQueue>[] = [
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "Type",
    header: "Type",
  },
  {
    accessorKey: "IsRegistered",
    header: () => <b>Is Registered</b>,
    cell: ({ row }) => <CellMaintening Queue={row.original} />,
  },
  {
    id: "actions",
    header: () => <b>Actions</b>,
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
