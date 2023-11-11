"use client";
import { RabbitMqCluster } from "@/types";
import { useTranslations } from "next-intl";
import React from "react";

function InfoCluster({ cluster }: { cluster: RabbitMqCluster }) {
  const t = useTranslations();
  return (
    <div className="py-6 flex flex-col space-y-2">
      <div>
        <h2 className="text-xl">Name</h2>
        <p className="text-sm text-muted-foreground">{cluster.name}</p>
      </div>
      <div>
        <h2 className="text-xl">Description</h2>
        <p className="text-sm text-muted-foreground">{cluster.description}</p>
      </div>
      <div>
        <h2 className="text-xl">Created At</h2>
        <p className="text-sm text-muted-foreground">
          {new Date(cluster.CreatedAt).toLocaleDateString()}
        </p>
      </div>
      <div className="flex gap-x-6 border w-fit p-2 rounded-md">
        <div className="bg-zinc-50 w-fit px-5 py-2 rounded-md hover:cursor-pointer">
          <h2 className="text-xl ">Host</h2>
          <p className="text-sm text-muted-foreground">{cluster.host}</p>
        </div>
        <div className="bg-zinc-50 w-fit px-5 py-2 rounded-md hover:cursor-pointer">
          <h2 className="text-xl">{t("Commons.User")}</h2>
          <p className="text-sm text-muted-foreground">{cluster.user}</p>
        </div>
        <div className="bg-zinc-50 w-fit px-5 py-2 rounded-md hover:cursor-pointer">
          <h2 className="text-xl">{t("Commons.Port")}</h2>
          <p className="text-sm text-muted-foreground">{cluster.port}</p>
        </div>
      </div>
    </div>
  );
}

export default InfoCluster;
