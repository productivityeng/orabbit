import React, { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { Copy, Edit, MoreHorizontal, Trash } from "lucide-react";
import toast from "react-hot-toast";
import { useParams, useRouter } from "next/navigation";
import { RabbitMqQueue } from "@/types";

interface CellActionProps {
  data: RabbitMqQueue;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [deleteModalOpen, setDeleteModalOpen] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);

  const onCopy = (id: string) => {
    navigator.clipboard.writeText(id);
    toast.success("Copied to clipboard");
  };

  const onDeleteStore = async () => {
    const toastId = toast.loading("Deleting Billboard...");
    try {
      setDeleteLoading(true);
      //await axios.delete(`/api/${params.storeId}/billboards/${data.id}`);
      toast.success("Billboard deleted!", { id: toastId });
      router.refresh();
    } catch (error) {
      toast.error(
        "Make sure you removed all categories using this billboard first",
        { id: toastId }
      );
      console.log(error);
    } finally {
      setDeleteLoading(false);
      setDeleteModalOpen(false);
    }
  };

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Button variant={"ghost"} className="w-8 h-8 p-0">
            <span className="sr-only">Open Menu</span>
            <MoreHorizontal className="h-4 w-4" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Actions</DropdownMenuLabel>
          <DropdownMenuItem
            onClick={() =>
              router.push(`/${params.storeId}/billboards/${data.ID}`)
            }
          >
            <Edit className="mr-2 h-4 w-4" /> Update
          </DropdownMenuItem>{" "}
          <DropdownMenuItem onClick={() => onCopy(data.ID.toString())}>
            <Copy className="mr-2 h-4 w-4" /> Copy Id
          </DropdownMenuItem>{" "}
          <DropdownMenuItem onClick={() => setDeleteModalOpen(true)}>
            <Trash className="mr-2 h-4 w-4" /> Delete
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
}

export default CellAction;
