"use client";

import { ColumnDef } from "@tanstack/react-table";
import VirtualHostCellAction from "./cell-action";
import { VirtualHosts } from "../../models/virtualhosts";
import CellVirtualHostMaintening from "./cell-maintening";

export const VirtualHostColumns: ColumnDef<VirtualHosts>[] = [
  {
    accessorKey: "Id",
    header: "Id",
  },
  {
    accessorKey: "Name",
    header: "Name",
  },
  {
    accessorKey: "Description",
    header: "Type",
  },
  {
    accessorKey: "IsRegistered",
    header: () => <b>Is Registered</b>,
    cell: ({ row }) => <CellVirtualHostMaintening VirtualHost={row.original} />,
  },
  {
    id: "actions",
    header: () => <b>Actions</b>,
    cell: ({ row }) => <VirtualHostCellAction data={row.original} />,
  },
];
