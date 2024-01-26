"use client";
import ImportClusterForm from "@/components/Modals/ImportCluster/ImportClusterModal";
import { useImportCluster } from "@/hooks/cluster-import";
import React, { useEffect, useState } from "react";

function ImportClusterProvider() {
  const [isMounted, setIsMounted] = useState(false);
  useEffect(() => {
    setIsMounted(true);
  }, []);

  if (!isMounted) return null;
  return (
    <>
      <ImportClusterForm />
    </>
  );
}

export default ImportClusterProvider;
