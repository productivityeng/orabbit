"use client";
import { cn } from "@/lib/utils";
import { LucideIcon } from "lucide-react";
import { useTranslations } from "next-intl";
import React from "react";
interface Props {
  Icon: LucideIcon;
  IconColor: string;
  BgIconColor: string;
  Titlei18Label: string;
  Descriptioni18Label: string;
}
function TopPage({
  Icon,
  IconColor,
  Titlei18Label,
  Descriptioni18Label,
  BgIconColor,
}: Props) {
  const t = useTranslations();
  return (
    <div className="flex gap-x-4 ">
      <div className={cn("aspect-square p-2 rounded-sm", BgIconColor)}>
        <Icon className={cn("w-10 h-10", IconColor)} />
      </div>
      <div>
        <div className="text-2xl">{t(Titlei18Label)}</div>
        <p className="text-muted-foreground text-sm">
          {t(Descriptioni18Label)}
        </p>
      </div>
    </div>
  );
}

export default TopPage;
