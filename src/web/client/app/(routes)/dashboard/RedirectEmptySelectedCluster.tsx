"use client";

import { useAppState } from "@/store/appstate";
import React, { useEffect } from "react";

function RedirectEmptySelectedCluster() {
  const { SetSelectedClusterId, SelectedClusterId } = useAppState();
  useEffect(() => {
    console.debug(
      "Executing useEffect from RedirectEmptySelectedCluster " +
        SelectedClusterId,
      SetSelectedClusterId
    );
    if (SelectedClusterId && SelectedClusterId > 0) {
      SetSelectedClusterId(undefined);
    }
  }, []);
  return <>{SelectedClusterId}</>;
}

export default RedirectEmptySelectedCluster;
