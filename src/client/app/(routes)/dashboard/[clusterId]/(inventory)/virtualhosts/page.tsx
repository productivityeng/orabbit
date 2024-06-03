"use client";
import {
  ImportVirtualHost,
  fetchVirtualHosts,
  removeVirtualHostAction,
  syncronizeVirtualHostAction,
} from "@/actions/virtualhost";
import React from "react";
import { VirtualHostsTable } from "./components/virtual-host-table/virtual-host-table";
import { RabbitMqVirtualHostColumnDef } from "./components/virtual-host-table/columns";
import { useQuery } from "@tanstack/react-query";
import { redirect, useParams } from "next/navigation";
import { VirtualTableContext } from "./components/virtual-host-table/virtualhost-table-context";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import toast from "react-hot-toast";
import { useTranslations } from "next-intl";
import _ from "lodash";
import { CreateLockerAction } from "@/actions/locker";

function VirtualHostPage() {
  const params = useParams() as { clusterId: string };
  const t = useTranslations("Dashboard.VirtualhostPage");

  const { data, refetch } = useQuery({
    queryKey: ["virtualhosts", params.clusterId],
    queryFn: async () => fetchVirtualHosts(Number(params.clusterId)),
  });

  if (params.clusterId === undefined) {
    redirect("/dashboard");
  }

  async function HandleImportVirtualHost(vhost: RabbitMqVirtualHost) {
    const toastId = toast.loading(t("Toast.ImportingVirtualHost"));
    try {
      await ImportVirtualHost(vhost.ClusterId, vhost.Name);
      toast.success(t("Toast.ImportingVirtualHostSuccess"), {
        id: toastId,
      });
      await refetch();
    } catch (error) {
      toast.error(t("Toast.ImportingVirtualHostError"), {
        id: toastId,
      });
    }
  }

  async function HandleRemoveVirtualHost(vhost: RabbitMqVirtualHost) {
    let toastId = toast.loading(t("Toast.RemovingVirtualHost"));
    try {
      let result = await removeVirtualHostAction(
        Number(vhost.ClusterId),
        vhost.Id
      );
      if (result.Result) {
        toast.success(t("Toast.VirtualHostSuccessRemoved"), { id: toastId });
        await refetch();
      } else {
        toast.error(t("Toast.VirtualHostFailToRemove"), {
          id: toastId,
        });
      }
    } catch (error) {
      toast.error(`Falha ao remover VirtualHost:${error}`, { id: toastId });
    }
  }

  async function HandleSyncronizeVirtualHost(vhost: RabbitMqVirtualHost) {
    let toastId = toast.loading(t("Toast.SyncronizingVirtualHost"));
    try {
      let result = await syncronizeVirtualHostAction(
        Number(vhost.ClusterId),
        vhost.Id
      );
      if (result.Result) {
        toast.success(t("Toast.VirtualHostSyncSuccess"), { id: toastId });
        await refetch();
      } else {
        toast.error(t("Toast.VirtualHostSyncFail"), {
          id: toastId,
        });
      }
    } catch (error) {
      toast.error(t("Toast.VirtualHostSyncFail"), {
        id: toastId,
      });
    }
  }

  async function HandleLockItem(
    virtualHost: RabbitMqVirtualHost,
    reason: string
  ) {
    let toastId = toast.loading(t("Toast.LockingVirtualHost"));
    try {
      let result = await CreateLockerAction(
        Number(params.clusterId),
        "virtualhost",
        virtualHost.Id,
        {
          reason: reason,
          responsible: "Victor",
        }
      );
      if (result.Result) {
        toast.success(t("Toast.LockingVirtualHost.Success"), { id: toastId });
        await refetch();
      } else {
        toast.error(t("Toast.LockingVirtualHost.Fail"), { id: toastId });
      }
    } catch (error) {
      toast.error(t("Toast.LockingVirtualHost.Fail"), { id: toastId });
    }
  }

  return (
    <VirtualTableContext.Provider
      value={{
        OnImportVirtualHostClick: HandleImportVirtualHost,
        OnRemoveTrackingFromVirtualHost: HandleRemoveVirtualHost,
        OnSyncronizeVirtualHost: HandleSyncronizeVirtualHost,
        HandleLockItem: HandleLockItem,
      }}
    >
      <main>
        <VirtualHostsTable
          columns={RabbitMqVirtualHostColumnDef}
          data={_.sortBy(data, (x) => x.Id) ?? []}
        />
      </main>
    </VirtualTableContext.Provider>
  );
}

export default VirtualHostPage;
