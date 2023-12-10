"use client";

import { ColumnDef } from "@tanstack/react-table";
import CellAction from "./cell-action";
import CellMaintening from "./cell-maintening";
import { RabbitMqUser } from "@/models/users";
import LockItem from "@/components/lock-item/lock-item";

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
    header: () => <b>Actions</b>,
    cell: ({ row }) => (
      <>
        {" "}
        <LockItem
          isLocked={false}
          lockType="User"
          artifactName={row.original.Username}
        />
        <CellAction data={row.original} />
      </>
    ),
  },
];
