import React from "react";
import { fetchVirtualHosts } from "./functions/fetch-virtualhosts";
import VirtualHostsClient from "./components/client/client-virtualhost";

async function VirtualHostPage({ params }: { params: { clusterId: number } }) {
  const vhosts = await fetchVirtualHosts(params.clusterId);
  return <VirtualHostsClient data={vhosts} />;
}

export default VirtualHostPage;
