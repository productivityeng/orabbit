import React, { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { Edit, Files, RefreshCw, SettingsIcon } from "lucide-react";
import { useParams, useRouter } from "next/navigation";
import { cn } from "@/lib/utils";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";

interface CellActionProps {
  data: RabbitMqVirtualHost;
}

function CellAction({ data }: CellActionProps) {
  const router = useRouter();
  const params = useParams();

  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const removeQueueFromCluster = async () => {};

  const importQueueFromCluster = async () => {};

  async function syncronizeQueue() {
    router.refresh();
  }

  return (
    <DropdownMenu onOpenChange={setIsMenuOpen}>
      <DropdownMenuTrigger asChild>
        <Button
          variant={"ghost"}
          className="w-8 h-8 p-0 focus-visible:ring-0  focus-visible:ring-offset-0"
        >
          <SettingsIcon
            className={cn("w-4 h-4 duration-200 ease-in-out ", {
              "text-rabbit": isMenuOpen,
            })}
          />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <DropdownMenuLabel>Actions</DropdownMenuLabel>
        {data.IsInDatabase && data.IsInCluster && (
          <DropdownMenuItem onClick={removeQueueFromCluster}>
            <Edit className="mr-2 h-4 w-4" /> Remove From Cluster
          </DropdownMenuItem>
        )}
        {data.IsInCluster && !data.IsInDatabase && (
          <DropdownMenuItem onClick={importQueueFromCluster}>
            <Files className="mr-2 h-4 w-4" /> Importar
          </DropdownMenuItem>
        )}
        {data.IsInDatabase && !data.IsInCluster && (
          <DropdownMenuItem onClick={syncronizeQueue}>
            <RefreshCw className="mr-2 h-4 w-4" /> Syncronize
          </DropdownMenuItem>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

export default CellAction;
