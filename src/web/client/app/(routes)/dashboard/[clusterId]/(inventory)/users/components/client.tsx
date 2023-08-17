"use client";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import React from "react";
import { UserColumn, columns } from "./columns";
import { DataTable } from "@/components/ui/data-table";
import Heading from "@/components/Heading/Heading";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";

interface UsersClientProps {
  data: UserColumn[];
}

function UsersClient({ data }: UsersClientProps) {
  const t = useTranslations();
  return (
    <div>
      <SimpleHeading
        title={t("UsersPage.TrackedUsers")}
        description={t("UsersPage.TrackedUsersDescription")}
      />
      <DataTable searchKey="username" columns={columns} data={data} />
    </div>
  );
}
export default UsersClient;
