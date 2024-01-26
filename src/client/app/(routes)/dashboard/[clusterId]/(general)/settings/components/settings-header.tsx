"use client";
import Heading from "@/components/Heading/Heading";
import React from "react";
import DeleteCluseter from "./delete-cluster";
import { Settings } from "lucide-react";

function SettingsHeader() {
  return (
    <Heading
      Icon={Settings}
      IconColor="text-zinc-500 "
      Titlei18Label="Commons.Settings"
      BgIconColor="bg-zinc-200/50"
      Descriptioni18Label="SettingsPage.TopDescription"
    >
      <DeleteCluseter />
    </Heading>
  );
}

export default SettingsHeader;
