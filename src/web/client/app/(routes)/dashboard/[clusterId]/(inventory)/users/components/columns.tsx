"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellAction from "./cell-action";

export type UserColumn = {
  id: string;
  username: string;
  passwordHash: string;
};

export const columns: ColumnDef<UserColumn>[] = [
  {
    accessorKey: "username",
    header: "Username",
  },
  {
    accessorKey: "passwordHash",
    header: "PasswordHash",
  },
  {
    id: "actions",
    cell: ({ row }) => <CellAction data={row.original} />,
  },
];
