"use client";
import { ImportVirtualHost, fetchVirtualHosts } from "@/actions/virtualhost";
import React from "react";
import { VirtualHostsTable } from "./components/virtual-host-table/exchange-table";
import { RabbitMqVirtualHostColumnDef } from "./components/virtual-host-table/columns";
import { useQuery } from "@tanstack/react-query";
import { redirect, useParams } from "next/navigation";
import { VirtualTableContext } from "./components/virtual-host-table/virtualhost-table-context";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import toast from "react-hot-toast";
import { useTranslations } from "next-intl";

function VirtualHostPage() {
  const params = useParams() as { clusterId: string };
  const t = useTranslations("Dashboard.VirtualhostPage");

  const { data, isLoading, refetch } = useQuery({
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

  return (
    <VirtualTableContext.Provider
      value={{
        OnImportVirtualHostClick: HandleImportVirtualHost,
      }}
    >
      <main>
        <VirtualHostsTable
          columns={RabbitMqVirtualHostColumnDef}
          data={data ?? []}
        />
      </main>
    </VirtualTableContext.Provider>
  );
}

export default VirtualHostPage;
