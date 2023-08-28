import React from "react";
import UserForm from "./components/queue-form";
import { fetchUser } from "@/actions/users";
import { fetchQueue } from "@/actions/queue";
import { FrontResponse } from "@/actions/common/frontresponse";
import { RabbitMqQueue } from "@/types";
import { Translation } from "next-i18next";

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
        <UserForm initialData={fetchedQueue?.Result ?? null} />
      )}
    </div>
  );
}

export default UserPage;
