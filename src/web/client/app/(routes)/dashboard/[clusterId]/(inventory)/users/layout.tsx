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
  const homeUserSlice = "/users";
  return (
    <div className="max-h-screen">
      <Heading
        Icon={UserIcon}
        IconColor="text-purple-500 "
        Titlei18Label="Commons.User"
        BgIconColor="bg-purple-200/50"
        Descriptioni18Label="UsersPage.TopDescription"
      >
        <Separator />
        <Link hidden={!pathname.endsWith(homeUserSlice)} href={"users/new"}>
          <Button size="sm">
            <Plus className="w-4 h-4 mr-2" /> New User
          </Button>
        </Link>
      </Heading>
      <Separator className="my-2" />
      {children}
    </div>
  );
}

export default UsersLayout;
