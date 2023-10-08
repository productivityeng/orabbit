import { fetchQeueusFromCluster } from "@/actions/queue";
import Heading from "@/components/Heading/Heading";
import { Mail } from "lucide-react";
import { Metadata } from "next";
import React from "react";
import QueueClient from "./components/client";

async function QueuesPage({ params }: { params: { clusterId: number } }) {
  let queuesFromCluster = await fetchQeueusFromCluster(params.clusterId);
  if (!queuesFromCluster.Result || queuesFromCluster.Result.length === 0) {
    return <p>No queues for this cluster yet!</p>;
  }
  return <QueueClient data={queuesFromCluster.Result!} />;
}

export default QueuesPage;
