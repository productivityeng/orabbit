import React from "react";
import QueueForm from "./components/queue-form";
import { fetchQueue } from "@/actions/queue";
import { FrontResponse } from "@/actions/common/frontresponse";
import { RabbitMqQueue } from "@/models/queues";

async function UserPage({
  params,
}: {
  params: { queueId: string; clusterId: number };
}) {
  let fetchedQueue: FrontResponse<RabbitMqQueue | null> | null = null;

  if (params.queueId != "new") {
    fetchedQueue = await fetchQueue(parseInt(params.queueId), params.clusterId);
  }

  return (
    <div>
      {fetchedQueue?.ErrorMessage != null ? (
        <p>Queue not found</p>
      ) : (
        <QueueForm initialData={fetchedQueue?.Result ?? null} />
      )}
    </div>
  );
}

export default UserPage;
