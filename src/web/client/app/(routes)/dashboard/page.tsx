import ImportClusterAction from "@/components/ImportCluster/ImportClusterAction";
import React from "react";
import RedirectEmptySelectedCluster from "./RedirectEmptySelectedCluster";

function DashboardHome() {
  return (
    <div className="h-full w-full flex justify-center items-center  ">
      <div className="w-[20vw]">
        <ImportClusterAction ImportRoute="/dashboard/new-cluster" />{" "}
        <RedirectEmptySelectedCluster />
      </div>
    </div>
  );
}

export default DashboardHome;
