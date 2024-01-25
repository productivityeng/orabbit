import { fetchRegisteredUsers, fetchUsersFromCluster } from "@/actions/users";
import _ from "lodash";
import React from "react";
import { UserTable } from "./components/user-table/user-table";
import { RabbitMqUserTableColumnsDef } from "./components/user-table/columns";

async function UsersPage({ params }: { params: { clusterId: number } }) {
  const users = await fetchUsersFromCluster(params.clusterId);
  if (!users || users.length == 0) {
    return <p>Error</p>;
  }

  return (
    <div className="flex flex-col pt-5">
      <UserTable
        data={_.sortBy(users, (user) => user.Id)}
        columns={RabbitMqUserTableColumnsDef}
      />
    </div>
  );
}

export default UsersPage;
