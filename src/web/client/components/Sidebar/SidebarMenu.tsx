"use client";
import {
  ArrowBigDown,
  CompassIcon,
  Mail,
  Router,
  ShieldClose,
  User,
} from "lucide-react";
import { useTranslations } from "next-intl";
import Link from "next/link";
import React from "react";
import { useAppState } from "@/store/appstate";
import { cn } from "@/lib/utils";

const menuItems = [
  {
    label: "Commons.User",
    icon: User,
    href: "users",
    iconColor: "text-purple-500",
  },
  {
    label: "Commons.Queue",
    icon: Mail,
    href: "queues",
    iconColor: "text-orange-500",
  },
  {
    label: "Commons.Exchange",
    icon: Router,
    href: "exchanges",
    iconColor: "text-stone-500",
  },
];

function MenuItems() {
  const t = useTranslations();
  const SelectedClusterId = useAppState((state) => state.SelectedClusterId);

  return (
    <>
      <p className="px-6 divide-y-8 text-sm text-muted-foreground">
        Clusters items
      </p>
      {menuItems.map((item) => (
        <Link
          key={item.href}
          href={`/dashboard/${SelectedClusterId}/${item.href}`}
        >
          <div className="flex group items-center justify-start   space-x-6 mx-5 px-2 py-1  hover:bg-slate-600/80  rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
            <div>
              <item.icon
                className={cn("transition duration-200", item.iconColor)}
              />
            </div>
            <p className="text-white text-base"> {t(item.label)}</p>
          </div>
        </Link>
      ))}
      <div className="space-y-2">
        <p className="px-6 divide-y-8 text-xs text-muted-foreground">
          Compliance
        </p>
        <div className="flex items-center group justify-start space-x-6 mx-5 px-2 py-1  hover:bg-slate-600/80 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <ShieldClose className="group-hover:text-rabbit text-red-500 transition duration-200" />
          <p className="text-white text-xs lg:text-sm"> Drift detection</p>
        </div>
        <div className="flex  group items-center justify-start space-x-6 mx-5 px-2 py-1  hover:bg-slate-600/50 rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out">
          <CompassIcon className="group-hover:text-rabbit text-emerald-500 transition duration-200" />
          <p className="text-white text-xs lg:text-sm"> Trail</p>
        </div>
      </div>
    </>
  );
}
function SidebarMenu() {
  const t = useTranslations();
  const SelectedClusterId = useAppState((state) => state.SelectedClusterId);

  return (
    <div className="space-y-5">
      <div className="space-y-2">
        <div className="flex flex-col space-y-4 ">
          {SelectedClusterId ? (
            <MenuItems />
          ) : (
            <div className="flex flex-col justify-center items-center">
              <div className="text-center bg-slate-600/20 mx-4 py-2 rounded-md w-[90%]">
                {t("Sidebar.ClusterSelect")}
              </div>{" "}
              <ArrowBigDown className="h-16 w-16" />
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default SidebarMenu;
