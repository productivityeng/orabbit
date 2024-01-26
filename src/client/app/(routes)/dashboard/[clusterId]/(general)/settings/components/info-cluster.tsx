"use client";
import { RabbitMqCluster } from "@/types";
import React from "react";

function InfoCluster({ cluster }: { cluster: RabbitMqCluster }) {
  return (
    <div className="py-6 flex flex-col space-y-2">
      <div>
        <h2 className="text-xl">Name</h2>
        <p className="text-sm text-muted-foreground">{cluster.Name}</p>
      </div>
      <div>
        <h2 className="text-xl">Description</h2>
        <p className="text-sm text-muted-foreground">{cluster.Description}</p>
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
          <p className="text-sm text-muted-foreground">{cluster.Host}</p>
        </div>
        <div className="bg-zinc-50 w-fit px-5 py-2 rounded-md hover:cursor-pointer">
          <h2 className="text-xl">Usuario</h2>
          <p className="text-sm text-muted-foreground">{cluster.User}</p>
        </div>
        <div className="bg-zinc-50 w-fit px-5 py-2 rounded-md hover:cursor-pointer">
          <h2 className="text-xl">Porta</h2>
          <p className="text-sm text-muted-foreground">{cluster.Port}</p>
        </div>
      </div>
    </div>
  );
}

export default InfoCluster;
