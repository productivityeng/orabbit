"use client";
import React, { useEffect } from "react";
import { SelectCluster } from "./SelectCluster";
import SidebarMenu from "./SidebarMenu";
import { RabbitMqCluster } from "@/types";
import { useAppState } from "@/store/appstate";

function Sidebar({ Clusters }: { Clusters: RabbitMqCluster[] }) {
  const {
    SetAvailableClusters,
    SetSelectedClusterId,
    SelectedClusterId,
    GetSelectedCluster,
  } = useAppState();
  useEffect(() => {
    SetAvailableClusters(Clusters);
  }, [SelectedClusterId]);
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
        <SelectCluster
          SelectedCluster={GetSelectedCluster()}
          Clusters={Clusters}
          SetSelectedClusterId={SetSelectedClusterId}
        />
      </div>
    </div>
  );
}

export default Sidebar;
