import { fetchRegisteredUsers, fetchUsersFromCluster } from "@/actions/users";
import React from "react";
import UsersClient from "./components/client";

async function UsersPage({ params }: { params: { clusterId: number } }) {
  const users = await fetchUsersFromCluster(params.clusterId);
  if (!users || users.length == 0) {
    return <p>Error</p>;
  }

  return (
    <div className="">
      <UsersClient data={users} />
    </div>
  );
}

export default UsersPage;
