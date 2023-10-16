import { fetchRegisteredUsers, fetchUsersFromCluster } from "@/actions/users";
import _ from "lodash";
import React from "react";
import UsersClient from "./components/client";

async function UsersPage({ params }: { params: { clusterId: number } }) {
  const users = await fetchUsersFromCluster(params.clusterId);
  if (!users || users.length == 0) {
    return <p>Error</p>;
  }

  return (
    <div className="">
      <UsersClient data={_.sortBy(users, (user) => user.Id)} />
    </div>
  );
}

export default UsersPage;
