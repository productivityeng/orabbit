"use client";

import Heading from "@/components/Heading/Heading";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { Plus, User, UserIcon } from "lucide-react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";

function UsersLayout({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const homeUserSlice = "/users";
  return (
    <div className="max-h-screen">
      <Heading
        Icon={UserIcon}
        IconColor="text-stone-500 "
        Titlei18Label="Exchanges"
        BgIconColor="bg-stone-200/50"
        Descriptioni18Label="
                            Uma exchange é um ponto de entrada 
                            para mensagens, onde os produtores as enviam e de onde os consumidores as recebem. Ela recebe mensagens de produtores e as encaminha para filas com base em um conjunto de regras de roteamento conhecidas como `Bindings`. 
                            "
      >
        <Separator />
        <Link hidden={!pathname.endsWith(homeUserSlice)} href={"users/new"}>
          <Button size="sm">
            <Plus className="w-4 h-4 mr-2" /> New User
          </Button>
        </Link>
      </Heading>
      <Separator className="my-2" />
      {children}
    </div>
  );
}

export default UsersLayout;
