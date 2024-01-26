"use client";

import { useAppState } from "@/hooks/cluster";
import React, { useEffect } from "react";

function RedirectEmptySelectedCluster() {
  const { SetSelectedClusterId, SelectedClusterId } = useAppState();
  useEffect(() => {
    if (SelectedClusterId && SelectedClusterId > 0) {
      SetSelectedClusterId(undefined);
    }
  }, []);
  return <>{SelectedClusterId}</>;
}

export default RedirectEmptySelectedCluster;
