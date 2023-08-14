"use client";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import React from "react";
import { UserColumn, columns } from "./columns";
import { DataTable } from "@/components/ui/data-table";
import Heading from "@/components/Heading/Heading";
import SimpleHeading from "@/components/Heading/SimpleHeading";

interface UsersClientProps {
  data: UserColumn[];
}

function UsersClient({ data }: UsersClientProps) {
  return (
    <div>
      <SimpleHeading
        title="Users watched"
        description="Users that are registed in ostern"
      />
      <DataTable searchKey="username" columns={columns} data={data} />
    </div>
  );
}
export default UsersClient;
