import React, { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import {
  Copy,
  Edit,
  MoreHorizontal,
  RefreshCcw,
  RefreshCw,
  SettingsIcon,
  Trash,
} from "lucide-react";
import toast from "react-hot-toast";
import { useParams, useRouter } from "next/navigation";
import { RabbitMqQueue } from "@/models/queues";
import { cn } from "@/lib/utils";

interface CellActionProps {
  data: RabbitMqQueue;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [deleteModalOpen, setDeleteModalOpen] = useState(false);
  const [deleteLoading, setDeleteLoading] = useState(false);
  const [isMenuOpen, setIsMenuOpen] = useState(false);

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
      <DropdownMenu onOpenChange={setIsMenuOpen}>
        <DropdownMenuTrigger asChild>
          <Button
            variant={"ghost"}
            className="w-8 h-8 p-0 focus-visible:ring-0  focus-visible:ring-offset-0"
          >
            <SettingsIcon
              className={cn("w-4 h-4 duration-200 ease-in-out ", {
                "text-rabbit w-8 h-8": isMenuOpen,
              })}
            />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent align="end">
          <DropdownMenuLabel>Actions</DropdownMenuLabel>
          {data.IsInDatabase && (
            <DropdownMenuItem
              onClick={() =>
                router.push(`/${params.storeId}/billboards/${data.ID}`)
              }
            >
              <Edit className="mr-2 h-4 w-4" /> Remove From Track
            </DropdownMenuItem>
          )}
          {data.IsInCluster && (
            <DropdownMenuItem onClick={() => onCopy(data.ID.toString())}>
              <Copy className="mr-2 h-4 w-4" /> Remove From Cluster
            </DropdownMenuItem>
          )}
          {data.IsInDatabase && !data.IsInCluster && (
            <DropdownMenuItem onClick={() => setDeleteModalOpen(true)}>
              <RefreshCw className="mr-2 h-4 w-4" /> Syncronize
            </DropdownMenuItem>
          )}
        </DropdownMenuContent>
      </DropdownMenu>
    </>
  );
}

export default CellAction;
