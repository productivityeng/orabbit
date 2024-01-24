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

type SelectClusterProps = {
  Clusters: RabbitMqCluster[];
};

export function SelectCluster({ Clusters }: Readonly<SelectClusterProps>) {
  const router = useRouter();
  const t = useTranslations("Sidebar");
  const [open, setOpen] = React.useState(false);
  const [value, setValue] = React.useState("");

  return (
    <Popover open={open} onOpenChange={setOpen}>
      <PopoverTrigger asChild>
        <Button
          variant="outline"
          aria-expanded={open}
          className="w-full justify-between text-center text-slate-400 bg-slate-700 border-0 hover:bg-rabbit hover:text-slate-100 duration-200 ease-in-out"
        >
          {value
            ? Clusters.find((cluster) => cluster.Name === value)?.Name
            : t("ClusterSelect") + "..."}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-full ">
        <Command>
          <CommandInput
            role="form"
            placeholder={`${t("SearchForCluster")}...`}
            autoComplete="off"
          />
          <CommandEmpty>{t("NoClusterFounded")}</CommandEmpty>
          <CommandGroup>
            {Clusters.map((cluster) => (
              <CommandItem
                key={cluster.Name}
                onSelect={(currentValue) => {
                  setValue(currentValue === value ? "" : currentValue);
                  setOpen(false);
                  router.push(`/dashboard/${cluster.Id}`);
                }}
              >
                <Check
                  className={cn(
                    "mr-2 h-4 w-4",
                    value === cluster.Name ? "opacity-100" : "opacity-0"
                  )}
                />
                {cluster.Name}
              </CommandItem>
            ))}
          </CommandGroup>
        </Command>
      </PopoverContent>
    </Popover>
  );
}
