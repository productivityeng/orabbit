import React, { useState } from "react";
import toast from "react-hot-toast";
import { useParams, useRouter } from "next/navigation";
import AlertModal from "@/components/Modals/alert-danger-modal";
import { Trash } from "lucide-react";
import { UserColumn } from "./columns";
import { RabbitMqUser } from "@/types";
import { deleteUserFromRabbit } from "@/actions/users";

interface CellActionProps {
  data: RabbitMqUser;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [deleteModalOpen, setDeleteModalOpen] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);

  const onDeleteUserHandler = async () => {
    const toastId = toast.loading("Deletando usuario...");
    try {
      await deleteUserFromRabbit(Number(params.clusterId), data.Id);

      toast.success("Usuario removido do cluster!", { id: toastId });
      router.refresh();
    } catch (error) {
      toast.error("Erro ao deletar usuario", { id: toastId });
      console.log(error);
    } finally {
      setDeleteLoading(false);
      setDeleteModalOpen(false);
    }
  };

  return (
    <>
      <AlertModal
        isOpen={deleteModalOpen}
        onClose={() => setDeleteModalOpen(false)}
        onConfirm={onDeleteUserHandler}
        loading={deleteLoading}
      />
      <Trash
        onClick={() => setDeleteModalOpen(true)}
        className="hover:cursor-pointer hover:text-red-500 duration-200 ease-in-out"
      />
    </>
  );
}

export default CellAction;
