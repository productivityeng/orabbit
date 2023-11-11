"use client";
import { useParams, useRouter } from "next/navigation";
import { dashboardPageParams } from "./page";
import { useEffect } from "react";
import { useAppState } from "@/hooks/cluster";

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const params = useParams() as unknown as dashboardPageParams;
  const { SetSelectedClusterId } = useAppState();
  useEffect(() => {
    SetSelectedClusterId(params.clusterId);
  }, []);
  return <>{children}</>;
}
