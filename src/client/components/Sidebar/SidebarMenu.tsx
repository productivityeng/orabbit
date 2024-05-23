"use client";
import {
  ArrowBigDown,
  CompassIcon,
  Home,
  Mail,
  Router,
  Settings2,
  ShieldClose,
  User,
} from "lucide-react";
import Link from "next/link";
import React from "react";
import { useAppState } from "@/hooks/cluster";
import { cn } from "@/lib/utils";
import { usePathname } from "next/navigation";
import { useTranslations } from "next-intl";

const menuItems = [
  {
    label: "Common.User",
    icon: User,
    href: "users",
    iconColor: "text-purple-500",
  },
  {
    label: "Common.Queue",
    icon: Mail,
    href: "queuesandstreams",
    iconColor: "text-orange-500",
  },
  {
    label: "Common.Exchange",
    icon: Router,
    href: "exchanges",
    iconColor: "text-stone-500",
  },
  {
    label: "Common.VirtualHost",
    icon: Home,
    href: "virtualhosts",
    iconColor: "text-sky-500",
  },
];

const complianceItems = [
  {
    label: "Common.Drift",
    icon: ShieldClose,
    href: "drift",
    iconColor: "text-red-500",
  },
  {
    label: "Common.Trail",
    icon: CompassIcon,
    href: "trail",
    iconColor: "text-green-500",
  },
];

const generalItems = [
  {
    label: "Common.Settings",
    icon: Settings2,
    href: "settings",
    iconColor: "text-zinc-500",
  },
];

function MenuItems() {
  const SelectedClusterId = useAppState((state) => state.SelectedClusterId);
  const t = useTranslations();
  const pathname = usePathname();

  return (
    <>
      <p className="px-6 divide-y-8 text-sm text-muted-foreground truncate">
        Clusters items
      </p>
      {menuItems.map((item) => (
        <Link
          key={item.href}
          href={`/dashboard/${SelectedClusterId}/${item.href}`}
        >
          <div
            className={cn(
              "flex group items-center justify-start   space-x-6 mx-5 p-2  hover:bg-slate-600/80  rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out",
              {
                "bg-slate-600": pathname.includes(item.href),
              }
            )}
          >
            <div>
              <item.icon
                className={cn("transition duration-200", item.iconColor)}
              />
            </div>
            <p className="text-white text-base truncate hidden lg:block">
              {" "}
              {t(item.label)}
            </p>
          </div>
        </Link>
      ))}
      <div className="space-y-2">
        <p className="px-6 divide-y-8 text-xs text-muted-foreground truncate">
          Compliance
        </p>
        {complianceItems.map((item) => (
          <Link
            key={item.href}
            href={`/dashboard/${SelectedClusterId}/${item.href}`}
          >
            <div
              className={cn(
                "flex group items-center justify-start   space-x-6 mx-5 p-2  hover:bg-slate-600/80  rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out",
                {
                  "bg-slate-600": pathname.includes(item.href),
                }
              )}
            >
              <div>
                <item.icon
                  className={cn("transition duration-200", item.iconColor)}
                />
              </div>
              <p className="text-white text-base truncate  hidden lg:block">
              {t(item.label)}
              </p>
            </div>
          </Link>
        ))}
      </div>
      <div className="space-y-2">
        <p className="px-6 divide-y-8 text-xs text-muted-foreground truncate">
          General
        </p>{" "}
        {generalItems.map((item) => (
          <Link
            key={item.href}
            href={`/dashboard/${SelectedClusterId}/${item.href}`}
          >
            <div
              className={cn(
                "flex group items-center justify-start   space-x-6 mx-5 p-2  hover:bg-slate-600/80  rounded-md hover:cursor-pointer hover:scale-105 duration-200 ease-in-out",
                {
                  "bg-slate-600": pathname.includes(item.href),
                }
              )}
            >
              <div>
                <item.icon
                  className={cn("transition duration-200", item.iconColor)}
                />
              </div>
              <p className="text-white text-base truncat hidden lg:block ">
                {" "}
                {t(item.label)}
              </p>
            </div>
          </Link>
        ))}
      </div>
    </>
  );
}
function SidebarMenu() {
  const SelectedClusterId = useAppState((state) => state.SelectedClusterId);

  return (
    <div className="space-y-5">
      <div className="space-y-2">
        <div className="flex flex-col space-y-4 ">
          {SelectedClusterId ? <MenuItems /> : null}
        </div>
      </div>
    </div>
  );
}

export default SidebarMenu;
