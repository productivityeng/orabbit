"use client";
import AlertModal from "@/components/Modals/alert-danger-modal";
import { useClusterSettings } from "@/hooks/cluster-settings";
import { useParams } from "next/navigation";
import React, { useState } from "react";
import { toast } from "react-hot-toast";

function DeleteCluseter() {
  const [loading, setLoading] = useState(false);
  const { isDeleteModalOpen, closeDeleteModal, deleteCluster } =
    useClusterSettings();
  const { brokerId } = useParams() as unknown as { brokerId: number };
  const handleDeleteCluster = async () => {
    setLoading(true);
    let toastId = toast.loading("Deleting cluster...");
    try {
      let result = await deleteCluster(brokerId);
      if (result) {
        toast.success("Cluster deleted!", { id: toastId });
        closeDeleteModal();
      } else {
        toast.error("Something went wrong :(", { id: toastId });
      }
    } catch (error) {
      console.error("[DELETE_CLUSTER]", error);
      toast.error("Something went wrong :(", { id: toastId });
    } finally {
      setLoading(false);
    }
  };

  return (
    <>
      <AlertModal
        isOpen={isDeleteModalOpen}
        loading={loading}
        onClose={closeDeleteModal}
        onConfirm={handleDeleteCluster}
      />
    </>
  );
}

export default DeleteCluseter;
