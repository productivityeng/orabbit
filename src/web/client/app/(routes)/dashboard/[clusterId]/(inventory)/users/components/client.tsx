"use client";
import React from "react";
import { columns } from "./columns";
import { DataTable } from "@/components/ui/data-table";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";
import { RabbitMqUser } from "@/types";

interface UsersClientProps {
  data: RabbitMqUser[];
}

function UsersClient({ data }: UsersClientProps) {
  const t = useTranslations();
  return (
    <div>
      <SimpleHeading
        title={t("UsersPage.TrackedUsers")}
        description={t("UsersPage.TrackedUsersDescription")}
      />
      <DataTable searchKey="Username" columns={columns} data={data} />
    </div>
  );
}
export default UsersClient;
