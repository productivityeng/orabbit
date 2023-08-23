"use client";
import React, { useEffect } from "react";
import { SelectCluster } from "../SelectCluster/SelectCluster";
import SidebarMenu from "./SidebarMenu";
import { RabbitMqCluster } from "@/types";
import { useAppState } from "@/hooks/cluster";
import Link from "next/link";
import Image from "next/image";
import RabbitLogo from "../../public/svg/rabbitmq-logo.svg";
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
    <div className="bg-slate-900 text-slate-500 h-screen w-full flex flex-col items-center justify-between pb-16 space-y-12">
      <div className="w-full flex flex-col space-y-8">
        <Link href="/dashboard">
          <div className="flex justify-center py-5 text-white items-center space-x-2 font-semibold text-2xl truncate">
            <Image alt="" className="w-8 h-8" src={RabbitLogo} />
            <p className="truncate">Ostern</p>
          </div>
        </Link>
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
