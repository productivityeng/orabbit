"use client";
import React, { useEffect, useState } from "react";
import { Unlock } from "lucide-react";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { Button } from "../ui/button";
import LockItemForm from "./lock-item-form";
import { LockerModel } from "@/actions/locker";
import { z } from "zod";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";

interface LockItem {
  Lockers?: LockerModel[];
  Disabled: boolean;
  Label: string;
  onLockItem: (data: z.infer<typeof LockItemFormSchema>) => Promise<void>;
}

function LockItem({ Disabled, Label, onLockItem }: LockItem) {
  const [isDialogOpen, setDialogOpen] = useState(false);

  const [isMounted, setIsMounted] = useState(false);
  useEffect(() => {
    setIsMounted(true);
  }, [isMounted]);
  if (!isMounted) return null;

  return (
    <Dialog open={isDialogOpen} onOpenChange={(open) => setDialogOpen(open)}>
      <DialogTrigger asChild>
        <Button
          disabled={Disabled}
          className="h-8 gap-2"
          size="sm"
          variant="alert"
          data-testid="lock-unlock-button"
        >
          <>
            <Unlock data-testid="lock-icon" className="w-4 h-4  " />
            Trancar
          </>
        </Button>
      </DialogTrigger>
      <DialogContent className="min-w-max">
        <DialogHeader>
          <DialogTitle>Trancar {Label}</DialogTitle>
          <DialogDescription>
            Impeça que alterações não combinadas sejam realizadas
          </DialogDescription>
        </DialogHeader>
        <LockItemForm
          onFormSubmit={async (data) => {
            await onLockItem(data);
            setDialogOpen(false);
          }}
        />
      </DialogContent>
    </Dialog>
  );
}

export default LockItem;
