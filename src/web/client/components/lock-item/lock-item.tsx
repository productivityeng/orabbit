"use client";
import React, { useEffect, useState } from "react";
import { LockIcon, Unlock } from "lucide-react";
import toast from "react-hot-toast";
import { useTranslations } from "next-intl";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import { Label } from "../ui/label";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import LockItemForm from "./lock-item-form";
import {
  CreateLockerAction,
  GetLocker,
  LockerModel,
  LockerType,
} from "@/actions/locker";
import { GetActiveLocker } from "@/lib/utils";
import { z } from "zod";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";
import { set } from "lodash";

interface LockItem {
  Lockers?: LockerModel[];
  Disabled: boolean;
  Label: string;
  onLockItem: (data: z.infer<typeof LockItemFormSchema>) => Promise<void>;
}

function LockItem({ Disabled, Label, onLockItem }: LockItem) {
  const t = useTranslations();
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
          <DialogDescription>{t("label-block-item")}</DialogDescription>
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
