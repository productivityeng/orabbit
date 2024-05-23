"use client";
import { fetchQeueusFromCluster, removeQueueFromClusterAction, syncronizeQueueAction } from "@/actions/queue";
import React from "react";
import _ from "lodash";
import SimpleHeading from "@/components/Heading/SimpleHeading";
import { QueueTable } from "./components/queue-table/queue-table";
import { Button } from "@/components/ui/button";
import { FileStack, RefreshCcwDot } from "lucide-react";
import { RabbitMqQueueColumn } from "./components/queue-table/columns";
import { useParams } from "next/navigation";
import { useQuery } from "@tanstack/react-query";
import { QueueTableContext } from "./components/queue-table/queue-table-context";
import { RabbitMqQueue } from "@/models/queues";
import toast from "react-hot-toast";
import { useTranslations } from "next-intl";

function QueuesPage() {

  const t = useTranslations("Dashboard.QueuesPage");
  
  const params = useParams();
  const {data,isLoading,refetch} = useQuery({
    queryKey: ["users", params.clusterId],
    queryFn: async () => fetchQeueusFromCluster(Number(params.clusterId)),
  })

  if (isLoading) {
    return <p>Loading...</p>;
  }

  if(!data){ 
    return <p>Erro ao carregar filas</p>
  }

  async function onSyncronizeQueueClick(queue:RabbitMqQueue){
    if (!queue) return;
    let toastId = toast.loading(<p>{t("Toast.Syncronizing")}</p>);
    try {
      let result = await syncronizeQueueAction({
        ClusterId: Number(queue.ClusterId),
        QueueId: queue.ID,
      });
      if (!result.Result) {
        toast.error(<p>Erro {result.ErrorMessage} ao sincronizar filas</p>, {
          id: toastId,
        });
        return;
      }
      toast.success(<p>Filas sincronizadas</p>, { id: toastId });
      await refetch();
    } catch (error) {
      toast.error(<p>{t("Toast.SyncronizingError")}</p>, { id: toastId });
    }
  };

  async  function onRemoveQueueClick(queue:RabbitMqQueue){
    if (!queue) return;
    let toastId = toast.loading(<p>{t("Toast.Removing")}</p>);
    try {
      let result = await removeQueueFromClusterAction(
        queue.ClusterId,
        queue.ID
      );
      if (!result.Result) {
        toast.error(<p>{t("Toast.RemoveFail")}</p>, {
          id: toastId,
        });
        return;
      }
      toast.success(<p>{t("Toast.RemovedSuccess")}</p>, { id: toastId });
      await refetch();
    } catch (error) {
      toast.error(<p>{t("Toast.RemoveFail")}</p>, { id: toastId });
    }
  };

  return (
    <div className="flex flex-col pt-5">
      <SimpleHeading
        title={t("Title")}
        description={t("Description")}
      />
      <QueueTableContext.Provider value={{
        onSyncronizeQueueClick,
        onRemoveQueueClick,
        ClusterId: Number(params.clusterId)
      }} >
      <QueueTable
        columns={RabbitMqQueueColumn}
        data={_.sortBy(data.Result, (queue) => queue.ID)}
      />
      </QueueTableContext.Provider>

    </div>
  );
}

export default QueuesPage;
