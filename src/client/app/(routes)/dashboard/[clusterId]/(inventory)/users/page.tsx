"use client";
import { SyncronizeUserAction, fetchUsersFromCluster, removeUserFromCluster } from "@/actions/users";
import _ from "lodash";
import React from "react";
import { UserTable } from "./components/user-table/user-table";
import { RabbitMqUserTableColumnsDef } from "./components/user-table/columns";
import { useQuery } from "@tanstack/react-query";
import { useParams, useRouter } from "next/navigation";
import { UserTableContext } from "./components/user-table/user-table-context";
import { RabbitMqUser } from "@/models/users";
import toast from "react-hot-toast";

function UsersPage() {
  const params = useParams();
  const router = useRouter();
  const {data,isLoading,refetch} = useQuery({
    queryKey: ["users", params.clusterId],
    queryFn: async () => fetchUsersFromCluster(Number(params.clusterId)),

  })

  if (isLoading) {
    return <p>Loading...</p>;
  }


  async function onSyncronizeUserClick(user:RabbitMqUser){
    
    const toastId = toast.loading(
      `Sincronizando usuarios`
    );

    try {
      let result = await SyncronizeUserAction(
        user.ClusterId,
        user.Id
      );
      if (result.Result) {
        toast.success(
          `Usuario ${user.Username} sincronizado com sucesso`,
          {
            id: toastId,
          }
        );
        await refetch();

      } else {
        toast.error(
          `Error ao sincronizar usuario ${user.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao sincronizar usuario ${user.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  async function onRemoveUserHandler (user:RabbitMqUser){
    const toastId = toast.loading(`Removendo usuario ${user.Username}`);
    try {

      let result = await removeUserFromCluster(
        user.ClusterId,
        user.Id
      );
      if (result.Result) {
        toast.success(`Usuario ${user.Username} removido com sucesso`, {
          id: toastId,
        });
        await refetch();
      } else {
        toast.error(
          `Error ao remover usuario ${user.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao remover usuario ${user.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };
  
  return ( <UserTableContext.Provider value={{
      onSyncronizeUser: onSyncronizeUserClick,
      onRemoveUser: onRemoveUserHandler
  }}> 
    <div className="flex flex-col pt-5">
      <UserTable
        data={_.sortBy(data, (user) => user.Id)}
        columns={RabbitMqUserTableColumnsDef}
      />
    </div>
    </UserTableContext.Provider>
  );
}

export default UsersPage;
