import { fetchQeueusFromCluster } from "@/actions/queue";
import React from "react";
import QueueClient from "./components/queue-table/client";
import _ from "lodash";

async function QueuesPage({ params }: { params: { clusterId: number } }) {
  let queuesFromCluster = await fetchQeueusFromCluster(params.clusterId);
  if (!queuesFromCluster.Result || queuesFromCluster.Result.length === 0) {
    return <p>No queues for this cluster yet!</p>;
  }
  return (
    <QueueClient
      data={_.sortBy(queuesFromCluster.Result, (queue) => queue.ID)}
    />
  );
}

export default QueuesPage;
