"use client";

import Heading from "@/components/Heading/Heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Mail, Plus, User, UserIcon } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

function UsersLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const homeQueueSlice = "/queuesandstreams";
  return (
    <div className="max-h-screen">
      <Heading
        Icon={Mail}
        IconColor="text-orange-500 "
        Titlei18Label="Commons.Queue"
        BgIconColor="bg-orange-200/50"
        Descriptioni18Label="QueuesPage.TopDescription"
      >
        <Separator />
        <Link
          hidden={!pathname.endsWith(homeQueueSlice)}
          href={"queuesandstreams/new"}
        >
          <Button size="sm">
            <Plus className="w-4 h-4 mr-2" /> New Queue
          </Button>
        </Link>
      </Heading>
      <Separator className="my-2" />
      {children}
    </div>
  );
}

export default UsersLayout;
