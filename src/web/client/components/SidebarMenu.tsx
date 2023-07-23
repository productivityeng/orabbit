"use client";
import Cog8ToothIcon from "@heroicons/react/24/outline/Cog8ToothIcon";
import { Mail, Router, User } from "lucide-react";
import { useTranslations } from "next-intl";
import Link from "next/link";
import React from "react";

const menuItems = [
  {
    label: "Commons.User",
    icon: User,
    href: "/users",
  },
  {
    label: "Commons.Queue",
    icon: Mail,
    href: "/queues",
  },
  {
    label: "Commons.Exchange",
    icon: Router,
    href: "/exchanges",
  },
];

function SidebarMenu() {
  const t = useTranslations();
  return (
    <div className="space-y-5">
      <div className="space-y-2">
        <p className="px-6 divide-y-8">Clusters items</p>
        <div className="flex flex-col space-y-4">
          {menuItems.map((item) => (
            <Link key={item.href} href={item.href}>
              <div className="flex group items-center justify-start  space-x-6 mx-5 px-2 py-1 bg-slate-600/20 hover:bg-rabbit rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
                <item.icon className="group-hover:text-white" />
                <p className="text-white text-lg"> {t(item.label)}</p>
              </div>
            </Link>
          ))}
        </div>
      </div>

      <div className="space-y-2">
        <p className="px-6 divide-y-8">Compliance</p>
        <div className="flex items-center justify-start space-x-6 mx-5 px-2 py-1 bg-slate-600/20 hover:bg-slate-600/80 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <Cog8ToothIcon className="h-6" />{" "}
          <p className="text-white text-lg"> Drift detection</p>
        </div>
      </div>
    </div>
  );
}

export default SidebarMenu;
