import { fetchQeueusFromCluster } from "@/actions/queue";
import React from "react";
import _ from "lodash";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { QueueTable } from "./components/queue-table/queue-table";
import { Button } from "@/components/ui/button";
import { FileStack, RefreshCcwDot } from "lucide-react";
import { RabbitMqQueueColumn } from "./components/queue-table/columns";
import { useTranslations } from "next-intl/dist/react-client";

async function QueuesPage({ params }: { params: { clusterId: number } }) {
  let queuesFromCluster = await fetchQeueusFromCluster(params.clusterId);
  if (!queuesFromCluster.Result || queuesFromCluster.Result.length === 0) {
    return <p>No queues for this cluster yet!</p>;
  }

  return (
    <div className="flex flex-col pt-5">
      <SimpleHeading
        title={"Filas rastreadas"}
        description={"Todas as filas visiveis no cluster"}
      />
      <QueueTable
        searchKey="name"
        columns={RabbitMqQueueColumn}
        data={_.sortBy(queuesFromCluster.Result, (queue) => queue.ID)}
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

export default QueuesPage;
