"use client";
import React from "react";
import { DataTable } from "@/components/ui/data-table";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";
import { RabbitMqQueue } from "@/types";
import { RabbitMqQueueColumn } from "./columns";

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
      <DataTable searchKey="name" columns={RabbitMqQueueColumn} data={data} />
    </div>
  );
}
export default QueueClient;
