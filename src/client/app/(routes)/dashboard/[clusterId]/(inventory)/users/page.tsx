"use client";
import { SyncronizeUserAction, fetchUsersFromCluster, importUserFromCluster, removeUserFromCluster } from "@/actions/users";
import _ from "lodash";
import React from "react";
import { UserTable } from "./components/user-table/user-table";
import { RabbitMqUserTableColumnsDef } from "./components/user-table/columns";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "next/navigation";
import { UserTableContext } from "./components/user-table/user-table-context";
import { RabbitMqUser } from "@/models/users";
import toast from "react-hot-toast";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";
import { z } from "zod";
import { CreateLockerAction, RemoveLockerAction } from "@/actions/locker";

function UsersPage() {
  const params = useParams();
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

  async function  onImportUserHandler(user:RabbitMqUser){
    const toastId = toast.loading(`Importando usuario ${user.Username}`);
    try {
      let result = await importUserFromCluster({
        ClusterId: user.ClusterId,
        Username: user.Username,
        Create: false,
      });
      if (result.Result) {
        toast.success(`Usuario ${user.Username} importado com sucesso`, {
          id: toastId,
        });
        await refetch();
      } else {
        toast.error(
          `Error ao importar usuario ${user.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao importar usuario ${user.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  async function onLockUserHandler (user: RabbitMqUser,data: z.infer<typeof LockItemFormSchema>){
    const toastId = toast.loading(`Bloqueando usuario ${user.Username}`);

    try {
      let result = await CreateLockerAction(
        user.ClusterId,
        "user",
        user.Id,
        {
          reason: data.reason,
          responsible: "Victor",
        }
      );
      if (result.Result) {
        toast.success(`Usuario ${user.Username} bloqueado com sucesso`, {
          id: toastId,
        });
        await refetch();
      } else {
        toast.error(
          `Error ao bloquear usuario ${user.Username} => ${result.ErrorMessage}`,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      toast.error(
        `Error ao bloquear usuario ${user.Username} => ${error}`,
        {
          id: toastId,
        }
      );
    }
  };

  async function onRemoveLockerHandler(user:RabbitMqUser,lockId: number){
    let toastId = toast.loading(
      `Removendo bloqueio da fila ${user.Username}...`
    );
    try {
      await RemoveLockerAction(user.ClusterId, "user", lockId);
      toast.success(`Bloqueio removido com sucesso`, { id: toastId });
      await refetch();
    } catch (error) {
      toast.error(`Erro ${JSON.stringify(error)} ao remover bloqueio`, {
        id: toastId,
      });
    }
  };
  
  return ( <UserTableContext.Provider value={{
      onSyncronizeUser: onSyncronizeUserClick,
      onRemoveUser: onRemoveUserHandler,
      onImportUser: onImportUserHandler,
      onLockUser: onLockUserHandler,
      onUnlockUser: onRemoveLockerHandler,
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
