import { fetchVirtualHosts } from "@/actions/virtualhost";
import React from "react";
import { VirtualHostsTable } from "./components/virtual-host-table/exchange-table";
import { RabbitMqVirtualHostColumnDef } from "./components/virtual-host-table/columns";

async function VirtualHostPage({ params }: { params: { clusterId: number } }) {
  const vhosts = await fetchVirtualHosts(params.clusterId);
  return (
    <main>
      <VirtualHostsTable columns={RabbitMqVirtualHostColumnDef} data={vhosts} />
    </main>
  );
}

export default VirtualHostPage;
