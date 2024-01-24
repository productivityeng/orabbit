"use client";
import React from "react";
import { DataTable } from "@/components/ui/data-table";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";
import { RabbitMqQueueColumn } from "./columns";
import { RabbitMqQueue } from "@/models/queues";
import { Button } from "@/components/ui/button";
import { FileStack, RefreshCcwDot } from "lucide-react";
import toast from "react-hot-toast";
import { syncronizeQueueAction } from "@/actions/queue";
import { useParams, useRouter } from "next/navigation";
import { QueueTable } from "./queue-table";

interface QueueClientProps {
  data: RabbitMqQueue[];
}

function QueueClient({ data }: QueueClientProps) {
  const t = useTranslations();

  return (
    <div>
      <SimpleHeading
        title={t("QueuePage.TrackedQueues")}
        description={t("QueuePage.TrackedQueuesDescription")}
      />
      <QueueTable
        searchKey="name"
        columns={RabbitMqQueueColumn}
        data={data}
        extraActions={
          <>
            <Button size="sm" className="h-8">
              <FileStack className="w-4 h-4 mr-2" /> Importar
            </Button>
            <Button size="sm" className="h-8">
              <RefreshCcwDot className="w-4 h-4 mr-2" /> Sincronizar
            </Button>
          </>
        }
      />
    </div>
  );
}
export default QueueClient;
