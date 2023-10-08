"use client";
import React from "react";
import { UserColumn } from "./columns";
import { DataTable } from "@/components/ui/data-table";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";
import { RabbitMqUser } from "@/types";
import { Button } from "@/components/ui/button";
import { FileStack } from "lucide-react";

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
      <DataTable
        searchKey="Username"
        columns={UserColumn}
        data={data}
        extraActions={
          <>
            <Button className="bg-rabbit">
              {" "}
              <FileStack className="w-4 h-4 mr-2" /> Mass import
            </Button>
          </>
        }
      />
    </div>
  );
}
export default UsersClient;
