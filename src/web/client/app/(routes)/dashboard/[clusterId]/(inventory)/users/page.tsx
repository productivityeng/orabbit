import { fetchUsersFromCluster } from "@/services/users";
import React from "react";
import UsersClient from "./components/client";
import { UserColumn } from "./components/columns";

async function UsersPage({ params }: { params: { clusterId: number } }) {
  const users = await fetchUsersFromCluster(params.clusterId);
  const formattedUser: UserColumn[] = users.result.map((user) => ({
    id: user.Id.toString(),
    username: user.Username,
    passwordHash: user.PasswordHash,
  }));

  return (
    <div className="">
      <UsersClient data={formattedUser} />
    </div>
  );
}

export default UsersPage;
