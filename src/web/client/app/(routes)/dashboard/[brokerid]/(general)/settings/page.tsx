"use client";
import TopPage from "@/components/TopPage/TopPage";
import { Button } from "@/components/ui/button";
import { Settings2, Trash2 } from "lucide-react";
import React from "react";
import DeleteCluseter from "./components/delete-cluster";
import { useClusterSettings } from "@/hooks/cluster-settings";

function SettingsPage() {
  const { openDeleteModal } = useClusterSettings();
  return (
    <main className="h-full w-full ">
      <TopPage
        Icon={Settings2}
        IconColor="text-zinc-500 "
        Titlei18Label="Commons.Settings"
        BgIconColor="bg-zinc-200/50"
        Descriptioni18Label="SettingsPage.TopDescription"
      >
        <Button
          size="icon"
          variant="destructive"
          className="my-10"
          onClick={() => openDeleteModal()}
        >
          <Trash2 />
        </Button>
      </TopPage>
    </main>
  );
}

export default SettingsPage;
