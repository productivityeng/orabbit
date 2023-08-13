"use client";
import Heading from "@/components/Heading/Heading";
import { Mail } from "lucide-react";
import React from "react";

function QueuesPage() {
  return (
    <main>
      <Heading
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
