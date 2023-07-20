import Sidebar from "@/components/Sidebar";
import { Metadata } from "next";
import React from "react";

export const metadata: Metadata = {
  title: "ORabbit | Dashboard",
  description: "Generated by create next app",
};

async function DashboardLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-row h-screen">
      <div className="w-[25vw] md:w-[13vw] h-screen">
        <Sidebar />
      </div>
      {children}
    </div>
  );
}

export default DashboardLayout;