import { RemoveLockerAction } from "@/actions/locker";
import { Button } from "@/components/ui/button";
import { UnlockItem } from "@/components/unlock-item/unlock-item";
import { GetActiveLocker } from "@/lib/utils";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import { UnlockIcon } from "lucide-react";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import toast from "react-hot-toast";

interface CellLockerProps {
  data: RabbitMqVirtualHost;
}
function CellLocker({ data }: CellLockerProps) {
  let activeLocker = GetActiveLocker(data.Lockers);
  const router = useRouter();
  const [isMounted, setIsMounted] = useState(false);

  useEffect(() => {
    setIsMounted(true);
  }, [isMounted]);

  const onRemoveLocker = async () => {
    let toastId = toast.loading(`Removendo bloqueio da fila ${data.Name}...`);
    try {
      await RemoveLockerAction(
        data.ClusterId,
        "virtualhost",
        activeLocker?.Id!
      );
      toast.success(`Bloqueio removido com sucesso`, { id: toastId });
    } catch (error) {
      toast.error(`Erro ${JSON.stringify(error)} ao remover bloqueio`, {
        id: toastId,
      });
    } finally {
      router.refresh();
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
