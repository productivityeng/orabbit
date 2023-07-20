"use client";
import TopPage from "@/components/TopPage/TopPage";
import { Mail } from "lucide-react";
import React from "react";

function QueuesPage() {
  return (
    <main>
      <TopPage
        Icon={Mail}
        IconColor="text-orange-500 "
        Titlei18Label="Commons.Queue"
        BgIconColor="bg-orange-200/50"
        Descriptioni18Label="QueuesPage.TopDescription"
      />
    </main>
  );
}

export default QueuesPage;
