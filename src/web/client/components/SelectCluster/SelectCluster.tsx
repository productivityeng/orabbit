"use client";

import * as React from "react";
import { Check, ChevronsUpDown } from "lucide-react";

import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "@/components/ui/command";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { useTranslations } from "next-intl";
import { RabbitMqCluster } from "@/types";
import { useRouter } from "next/navigation";
import { CommandList } from "cmdk";

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
  const t = useTranslations("Sidebar");
  const [open, setOpen] = React.useState(false);

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          role="combobox"
          aria-expanded={open}
          className="w-[300px] justify-between text-center text-slate-400 bg-slate-700 border-0 hover:bg-rabbit hover:text-slate-100 duration-200 ease-in-out"
        >
          {SelectedCluster?.name ?? t("ClusterSelect") + "..."}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-[300px] p-0">
        <Command className="w-full">
          <CommandList>
            <CommandInput
              role="commandInput"
              placeholder={`${t("SearchForCluster")}...`}
            />
            <CommandEmpty>{t("NoClusterFounded")}</CommandEmpty>
            <CommandGroup className="w-full ">
              {Clusters.map((cluster) => (
                <CommandItem
                  key={"idx" + cluster.Id}
                  onSelect={(currentValue) => {
                    if (SelectedCluster?.Id == cluster.Id) {
                      router.push(`/dashboard`);
                      SetSelectedClusterId(undefined);
                    } else {
                      SetSelectedClusterId(cluster.Id);
                      router.push(`/dashboard/${cluster.Id}`);
                    }
                    setOpen(false);
                  }}
                >
                  <Check
                    className={cn("mr-2 h-4 w-4 opacity-0 ", {
                      "opacity-100":
                        SelectedCluster && SelectedCluster.Id === cluster.Id,
                    })}
                  />
                  {cluster.name}
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
