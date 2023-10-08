"use client";
import React from "react";
import { DataTable } from "@/components/ui/data-table";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";
import { RabbitMqQueueColumn } from "./columns";
import { RabbitMqQueue } from "@/models/queues";
import { Button } from "@/components/ui/button";
import { FileStack, RefreshCcwDot } from "lucide-react";

interface QueueClientProps {
  data: RabbitMqQueue[];
}

function QueueClient({ data }: QueueClientProps) {
  const t = useTranslations();
  return (
    <div>
      <SimpleHeading
        title={t("QueuePage.TrackedQueues")}
        description={t("QueuePage.rackedQueuesDescription")}
      />
      <DataTable
        searchKey="name"
        columns={RabbitMqQueueColumn}
        data={data}
        extraActions={
          <>
            <Button size="sm">
              {" "}
              <FileStack className="w-4 h-4 mr-2" /> Mass import
            </Button>
            <Button size="sm">
              {" "}
              <RefreshCcwDot className="w-4 h-4 mr-2" /> Mass syncronize
            </Button>
          </>
        }
      />
    </div>
  );
}
export default QueueClient;
