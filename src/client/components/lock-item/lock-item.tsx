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
import { useTranslations } from "next-intl";

interface LockItem {
  Lockers?: LockerModel[];
  Disabled: boolean;
  Label: string;
  onLockItem: (data: z.infer<typeof LockItemFormSchema>) => Promise<void>;
}

function LockItem({ Disabled, Label, onLockItem }: LockItem) {
  const [isDialogOpen, setDialogOpen] = useState(false);

  const [isMounted, setIsMounted] = useState(false);
  const t  = useTranslations("Common.Component")
  useEffect(() => {
    setIsMounted(true);
  }, [isMounted]);
  if (!isMounted) return null;

  return (
    <Dialog  open={isDialogOpen} onOpenChange={(open) => setDialogOpen(open)}>
      <DialogTrigger asChild>
        <Button
          disabled={Disabled}
          className="h-8 gap-2"
          size="sm"
          variant="alert"
          data-testid="lock-unlock-button"
        >
          <>
            <Unlock data-testid="lock-button-id" className="w-4 h-4  " />
            {t("Lock")}
          </>
        </Button>
      </DialogTrigger>
      <DialogContent data-testid="lock-item-dialog" className="min-w-max">
        <DialogHeader>
          <DialogTitle>{t("Lock")} {Label}</DialogTitle>
          <DialogDescription>
            {t("LockItemDialogDescription")}
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
