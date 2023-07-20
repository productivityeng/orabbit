"use client";
import React from "react";
import { useParams } from "next/navigation";

type dashboardPageParams = {
  brokerid: number;
};

function DashboardPage() {
  const params = useParams() as unknown as dashboardPageParams;

  return <div>Dashboard {params.brokerid}</div>;
}

export default DashboardPage;
