"use client";

import Heading from "@/components/Heading/Heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus, User, UserIcon } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

function UsersLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();

  const homeUserSlice = "/virtualhosts";
  return (
    <div className="max-h-screen">
      <Heading
        Icon={UserIcon}
        IconColor="text-sky-500 "
        Titlei18Label="Dashboard.VirtualhostPage.Title"
        BgIconColor="bg-sky-200/50"
        Descriptioni18Label="Dashboard.VirtualhostPage.Description"
      >
        <Separator />
        <Link
          hidden={!pathname.endsWith(homeUserSlice)}
          href={"virtualhosts/new"}
        >
          <Button size="sm">
            <Plus className="w-4 h-4 mr-2" /> Adicionar novo
          </Button>
        </Link>
      </Heading>
      <Separator className="my-2" />
      {children}
    </div>
  );
}

export default UsersLayout;
