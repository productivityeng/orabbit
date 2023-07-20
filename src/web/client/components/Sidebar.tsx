import React from "react";
import { SelectCluster } from "./SelectCluster";
import { fetchAllClusters } from "@/services/cluster";
import SidebarMenu from "./SidebarMenu";

async function Sidebar() {
  const cluster = await fetchAllClusters();
  return (
    <div className="bg-slate-900 text-slate-500 h-full w-full flex flex-col items-center justify-between pb-16 pt-12">
      <div className="w-full">
        <div className="text-center text-2xl text-white pb-8">
          Welcome to sidebar
        </div>
        <div>
          <SidebarMenu />
        </div>
      </div>
      <div className="w-full px-5">
        <SelectCluster Clusters={cluster.result} />
      </div>
    </div>
  );
}

export default Sidebar;
