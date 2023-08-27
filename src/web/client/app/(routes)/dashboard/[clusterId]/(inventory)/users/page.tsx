import { fetchRegisteredUsers, fetchUsersFromCluster } from "@/actions/users";
import React from "react";
import UsersClient from "./components/client";

async function UsersPage({ params }: { params: { clusterId: number } }) {
  const users = await fetchUsersFromCluster(params.clusterId);

  return (
    <div className="">
      <UsersClient
        data={users.sort((user1, user2) =>
          user1.Username > user2.Username ? 1 : -1
        )}
      />
    </div>
  );
}

export default UsersPage;
