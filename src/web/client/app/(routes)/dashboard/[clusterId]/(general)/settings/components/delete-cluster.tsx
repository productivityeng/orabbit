"use client";
import AlertModal from "@/components/Modals/alert-danger-modal";
import { useClusterSettings } from "@/hooks/cluster-settings";
import { useParams, useRouter } from "next/navigation";
import React, { useState } from "react";
import { toast } from "react-hot-toast";

function DeleteCluseter() {
  const [loading, setLoading] = useState(false);
  const { isDeleteModalOpen, closeDeleteModal, deleteCluster } =
    useClusterSettings();
  const { brokerid } = useParams() as unknown as { brokerid: number };
  const router = useRouter();

  const handleDeleteCluster = async () => {
    setLoading(true);
    let toastId = toast.loading("Deleting cluster...");
    try {
      console.log("DELETING CLUSTER", brokerid);
      let result = await deleteCluster(brokerid);
      if (result) {
        toast.success("Cluster deleted!", { id: toastId });
        closeDeleteModal();
        router.push("/");
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
