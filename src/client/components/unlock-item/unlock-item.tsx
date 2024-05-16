import { LockerModel, LockerType } from "@/actions/locker";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTrigger,
} from "../ui/alert-dialog";
import { Button } from "../ui/button";
import { LockIcon, UnlockIcon, XIcon } from "lucide-react";

interface UnlockItemProps {
  Locker: LockerModel | null;
  onRemoveLocker: () => Promise<void>;
}
export function UnlockItem({ Locker, onRemoveLocker }: UnlockItemProps) {
  return (
    <AlertDialog >
      <AlertDialogTrigger>
        <Button
          size="sm"
          data-testid="unlock-icon-button"
          variant={Locker ? "destructive" : "outline"}
          className=" items-center justify-center"
        >
          <UnlockIcon className="w-4 h-4 " />
        </Button>
      </AlertDialogTrigger>
      <AlertDialogContent data-testid="unlock-dialog">
        <AlertDialogHeader>Locked</AlertDialogHeader>
        <AlertDialogDescription>
          <span>
            The user <b>{Locker?.UserResponsibleEmail}</b> has locked this queue
            at{" "}
            {new Date(Locker!.CreatedAt).toLocaleDateString("pt-BR", {
              hour: "2-digit",
              minute: "2-digit",
            })}
            &nbsp; for the reason: &nbsp;
            <span className="font-medium">{Locker?.Reason}</span>
          </span>
        </AlertDialogDescription>
        <AlertDialogFooter>
          <AlertDialogCancel className="flex h-9 gap-x-2">
            <XIcon className="w-4 h-4" /> Fechar
          </AlertDialogCancel>
          <AlertDialogAction
            data-testid="unlock-action-button"
            onClick={onRemoveLocker}
            className="flex gap-x-2 h-9"
          >
            <UnlockIcon className="w-4 h-4" /> Desbloquear
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
