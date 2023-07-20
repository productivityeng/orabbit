import Cog8ToothIcon from "@heroicons/react/24/outline/Cog8ToothIcon";
import React from "react";

function SidebarMenu() {
  return (
    <div className="space-y-5">
      <div className="space-y-2">
        <p className="px-6 divide-y-8">Clusters items</p>
        <div className="flex items-center justify-start space-x-6 mx-5 px-2 py-1 bg-slate-600/20 hover:bg-slate-600/80 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <Cog8ToothIcon className="h-6" />{" "}
          <p className="text-white text-lg"> Users</p>
        </div>
        <div className="group flex items-center justify-start space-x-6 mx-5 px-2 py-1 bg-slate-600/20 hover:bg-rabbit/100 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <Cog8ToothIcon className="h-6 hover:group:text-white" />{" "}
          <p className="text-white text-lg"> Queues</p>
        </div>
        <div className="flex items-center justify-start space-x-6 mx-5 px-2 py-1 bg-slate-600/20 hover:bg-slate-600/80 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <Cog8ToothIcon className="h-6" />{" "}
          <p className="text-white text-lg"> Exchanges</p>
        </div>
      </div>

      <div className="space-y-2">
        <p className="px-6 divide-y-8">Compliance</p>
        <div className="flex items-center justify-start space-x-6 mx-5 px-2 py-1 bg-slate-600/20 hover:bg-slate-600/80 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <Cog8ToothIcon className="h-6" />{" "}
          <p className="text-white text-lg"> Drift detection</p>
        </div>
      </div>
    </div>
  );
}

export default SidebarMenu;
