import { LockerModel, RemoveLockerAction } from "@/actions/locker";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import { UnlockItem } from "@/components/unlock-item/unlock-item";
import { GetActiveLocker } from "@/lib/utils";
import { RabbitMqQueue } from "@/models/queues";
import { RabbitMqUser } from "@/models/users";
import { LockIcon, UnlockIcon, XIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { use, useEffect, useState } from "react";
import toast from "react-hot-toast";

interface CellLockerProps {
  User: RabbitMqUser;
}
function CellLocker({ User }: CellLockerProps) {
  let activeLocker = GetActiveLocker(User.Lockers);
  const router = useRouter();
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    setIsMounted(true);
  }, [isMounted]);

  const onRemoveLocker = async () => {
    let toastId = toast.loading(
      `Removendo bloqueio da fila ${User.Username}...`
    );
    try {
      await RemoveLockerAction(User.ClusterId, "user", activeLocker?.Id!);
      toast.success(`Bloqueio removido com sucesso`, { id: toastId });
      router.refresh();
    } catch (error) {
      toast.error(`Erro ${JSON.stringify(error)} ao remover bloqueio`, {
        id: toastId,
      });
    }
  };

  if (!isMounted) return null;

  if (activeLocker) {
    return <UnlockItem Locker={activeLocker} onRemoveLocker={onRemoveLocker} />;
  } else {
    return (
      <Button
        size="sm"
        variant="outline"
        className=" items-center justify-center"
      >
        <UnlockIcon className="w-4 h-4" />
      </Button>
    );
  }
}

export default CellLocker;
