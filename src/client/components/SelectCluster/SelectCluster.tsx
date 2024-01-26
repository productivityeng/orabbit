"use client";

import * as React from "react";
import { Check, ChevronsUpDown, PlusCircle } from "lucide-react";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
  CommandSeparator,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { RabbitMqCluster } from "@/types";
import { useRouter } from "next/navigation";
import { CommandList } from "cmdk";
import { useImportCluster } from "@/hooks/cluster-import";

type SelectClusterProps = {
  Clusters: RabbitMqCluster[];
  SelectedCluster: RabbitMqCluster | undefined;
  SetSelectedClusterId: (clusterId: number | undefined) => void;
};

export function SelectCluster({
  Clusters,
  SetSelectedClusterId,
  SelectedCluster,
}: SelectClusterProps) {
  const router = useRouter();
  const { openModal } = useImportCluster();
  const [open, setOpen] = React.useState(false);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-full justify-between text-center text-slate-400 bg-slate-700 border-0 hover:bg-rabbit hover:text-slate-100 duration-200 ease-in-out"
        >
          <p className="truncate">
            {SelectedCluster?.Name ?? "Não encontrado " + "..."}
          </p>
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full p-0 popover-content-width-same-as-its-trigger">
        <Command className="w-full">
          <CommandList>
            <CommandInput
              role="commandInput"
              placeholder={`Pesquise um cluster...`}
            />
            <CommandEmpty>Nenum cluster encontrado</CommandEmpty>
            <CommandGroup className="w-full ">
              {Clusters?.map((cluster) => (
                <CommandItem
                  key={"idx" + cluster.Id}
                  onSelect={(currentValue) => {
                    SetSelectedClusterId(cluster.Id);
                    router.push(`/dashboard/${cluster.Id}`);
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn("mr-2 h-4 w-4 opacity-0 ", {
                      "opacity-100":
                        SelectedCluster && SelectedCluster.Id === cluster.Id,
                    })}
                  />
                  <p className="truncate"> {cluster.Name}</p>
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
          <CommandSeparator />
          <CommandList>
            <CommandGroup>
              <CommandItem className=" flex space-x-4" onSelect={openModal}>
                <div className=" w-full flex space-x-2 hover:cursor-pointer">
                  <PlusCircle className="mr2 h-5 w-5" />
                  <span>Importar Cluster</span>
                </div>
              </CommandItem>
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
