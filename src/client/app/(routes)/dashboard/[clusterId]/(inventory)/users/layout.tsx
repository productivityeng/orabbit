"use client";

import Heading from "@/components/Heading/Heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus, User, UserIcon } from "lucide-react";
import { useTranslations } from "next-intl";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

function UsersLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const homeUserSlice = "/users";
  const t = useTranslations("Dashboard.UsersPage");
  return (
    <div className="max-h-screen">
      <Heading
        Icon={UserIcon}
        IconColor="text-purple-500 "
        Titlei18Label={t("Title")}
        BgIconColor="bg-purple-200/50"
        Descriptioni18Label={t("SubTitle")}
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
