"use client";

import { useParams } from "next/navigation";
import { dashboardPageParams } from "./page";
import { useEffect } from "react";
import { useAppState } from "@/hooks/appstate";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const { SetSelectedClusterId } = useAppState();

  const params = useParams() as unknown as dashboardPageParams;
  useEffect(() => {
    SetSelectedClusterId(params.brokerid);
  }, [params.brokerid]);
  return <>{children}</>;
}
