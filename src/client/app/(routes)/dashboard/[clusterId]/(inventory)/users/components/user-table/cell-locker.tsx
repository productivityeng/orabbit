import { Button } from "@/components/ui/button";
import { UnlockItem } from "@/components/unlock-item/unlock-item";
import { GetActiveLocker } from "@/lib/utils";
import { RabbitMqUser } from "@/models/users";
import { UnlockIcon } from "lucide-react";
import { useContext, useEffect, useState } from "react";
import { UserTableContext } from "./user-table-context";

interface CellLockerProps {
  User: RabbitMqUser;
}
function CellLocker({ User }: CellLockerProps) {
  let activeLocker = GetActiveLocker(User.Lockers);
  const [isMounted, setIsMounted] = useState(false);
  const {onUnlockUser} = useContext(UserTableContext);

  useEffect(() => {
    setIsMounted(true);
  }, [isMounted]);

  if (!isMounted) return null;

  if (activeLocker) {
    return <UnlockItem Locker={activeLocker} onRemoveLocker={async () => await onUnlockUser?.(User, activeLocker?.Id!)} />;
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
