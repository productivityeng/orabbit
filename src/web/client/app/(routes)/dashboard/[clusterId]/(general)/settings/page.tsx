import React from "react";

import { Separator } from "@/components/ui/separator";
import SettingsHeader from "./components/settings-header";
import { fetchCluster } from "@/services/cluster";
import InfoCluster from "./components/info-cluster";

async function SettingsPage({ params }: { params: { clusterId: number } }) {
  const cluster = await fetchCluster(params.clusterId);
  return (
    <main className="h-full w-full ">
      <SettingsHeader />
      <Separator />
      <InfoCluster cluster={cluster} />
    </main>
  );
}

export default SettingsPage;
