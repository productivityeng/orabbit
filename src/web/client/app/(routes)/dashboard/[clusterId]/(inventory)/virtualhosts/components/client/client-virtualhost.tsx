"use client";
import React from "react";
import { DataTable } from "@/components/ui/data-table";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { useTranslations } from "next-intl";
import { VirtualHostColumns } from "./columns-client-virtualhost";
import { RabbitMqQueue } from "@/models/queues";
import { Button } from "@/components/ui/button";
import { FileStack, RefreshCcwDot } from "lucide-react";
import toast from "react-hot-toast";
import { syncronizeQueueAction } from "@/actions/queue";
import { useParams, useRouter } from "next/navigation";
import { VirtualHosts } from "../../models/virtualhosts";

interface VirtualHostsClientProps {
  data: VirtualHosts[];
}

function VirtualHostsClient({ data }: VirtualHostsClientProps) {
  const t = useTranslations();
  const params = useParams();
  const router = useRouter();

  const massSyncronizeQueue = async () => {
    // toast.success(<p>Filas sincronizadas</p>, {
    //   id: toastId,
    // });
    // router.refresh();
  };

  return (
    <div>
      <SimpleHeading
        title={t("QueuePage.TrackedQueues")}
        description={t("QueuePage.TrackedQueuesDescription")}
      />
      <DataTable
        searchKey="name"
        columns={VirtualHostColumns}
        data={data}
        extraActions={
          <>
            <Button size="sm">
              {" "}
              <FileStack className="w-4 h-4 mr-2" /> Mass import
            </Button>
            <Button size="sm" onClick={massSyncronizeQueue}>
              {" "}
              <RefreshCcwDot className="w-4 h-4 mr-2" /> Mass syncronize
            </Button>
          </>
        }
      />
    </div>
  );
}
export default VirtualHostsClient;
