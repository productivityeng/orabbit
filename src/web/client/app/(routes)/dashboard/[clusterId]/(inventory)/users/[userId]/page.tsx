import React from "react";
import UserForm from "./components/user-form";
import { fetchUser } from "@/services/users";

async function UserPage({
  params,
}: {
  params: { userId: string; clusterId: number };
}) {
  let existingUserResponse = null;

  if (params.userId != "new") {
    existingUserResponse = await fetchUser(
      parseInt(params.userId),
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
