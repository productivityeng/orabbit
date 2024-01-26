"use client";
import React, { useEffect } from "react";
import { useParams } from "next/navigation";
import { useAppState } from "@/hooks/cluster";

export type dashboardPageParams = {
  clusterId: number;
};

function DashboardPage() {
  const params = useParams() as unknown as dashboardPageParams;
  const { SetSelectedClusterId } = useAppState();
  useEffect(() => {
    SetSelectedClusterId(params.clusterId);
  }, []);
  return <div>Dashboard {params.clusterId}</div>;
}

export default DashboardPage;
