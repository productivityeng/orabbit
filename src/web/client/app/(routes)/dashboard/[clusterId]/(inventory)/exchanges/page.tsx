"use client";
import Heading from "@/components/Heading/Heading";
import { Router } from "lucide-react";
import React from "react";

function ExchangesPage() {
  return (
    <main>
      <Heading
        Icon={Router}
        IconColor="text-stone-500 "
        Titlei18Label="Commons.Exchange"
        BgIconColor="bg-stone-200/50"
        Descriptioni18Label="ExchangePage.TopDescription"
      />
    </main>
  );
}

export default ExchangesPage;
