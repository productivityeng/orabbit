import React from "react";
import UserForm from "./components/user-form";
import { fetchUser } from "@/services/users";

async function UserPage({
  params,
}: {
  params: { userId: number; clusterId: number };
}) {
  let existingUserResponse = await fetchUser(params.userId, params.clusterId);

  return (
    <div>
      <UserForm initialData={existingUserResponse.Result} />
    </div>
  );
}

export default UserPage;
