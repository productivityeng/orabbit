import React from "react";
import UserForm from "./components/queue-form";
import { fetchUser } from "@/actions/users";

async function UserPage({
  params,
}: {
  params: { queueId: string; clusterId: number };
}) {
  let existingUserResponse = null;

  if (params.queueId != "new") {
    existingUserResponse = await fetchUser(
      parseInt(params.queueId),
      params.clusterId
    );
  }

  return (
    <div>
      <UserForm initialData={existingUserResponse?.Result ?? null} />
    </div>
  );
}

export default UserPage;
