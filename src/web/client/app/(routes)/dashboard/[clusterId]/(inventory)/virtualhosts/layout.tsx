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
  const t = useTranslations();

  const homeUserSlice = "/virtualhosts";
  return (
    <div className="max-h-screen">
      <Heading
        Icon={UserIcon}
        IconColor="text-sky-500 "
        Titlei18Label="Commons.VirtualHost"
        BgIconColor="bg-sky-200/50"
        Descriptioni18Label="VirtualHostsPage.TopDescription"
      >
        <Separator />
        <Link
          hidden={!pathname.endsWith(homeUserSlice)}
          href={"virtualhosts/new"}
        >
          <Button size="sm">
            <Plus className="w-4 h-4 mr-2" /> {t("VirtualHostsPage.CreateNew")}
          </Button>
        </Link>
      </Heading>
      <Separator className="my-2" />
      {children}
    </div>
  );
}

export default UsersLayout;
