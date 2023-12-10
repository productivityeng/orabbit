"use client";

import React from "react";
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

type LockTypes = "User" | "VirtualHost" | "Queue" | "Exchange";
interface LockItem {
  isLocked: boolean;
  lockType: LockTypes;
  artifactName: string;
}

function LockItem({ isLocked, lockType, artifactName }: LockItem) {
  const t = useTranslations();

  const lockItem = async () => {
    toast.success(t("lock-success"));
  };

  const unlockItem = async () => {
    toast.success(t("unlock-success"));
  };

  return (
    <Dialog>
      <DialogTrigger asChild>
        <Button variant="ghost" data-testid="lock-unlock-button">
          {" "}
          {isLocked ? (
            <Unlock
              data-testid="unlock-icon"
              className="w-4 h-4  text-yellow-800"
            />
          ) : (
            <LockIcon
              data-testid="lock-icon"
              className="w-4 h-4  text-yellow-800"
            />
          )}
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>
            Bloquear {t(`label-${lockType}`)} {artifactName}
          </DialogTitle>
          <DialogDescription>{t("label-block-item")}</DialogDescription>
        </DialogHeader>
        <LockItemForm onFormSubmit={(data) => lockItem()} />
      </DialogContent>
    </Dialog>
  );
}

export default LockItem;
