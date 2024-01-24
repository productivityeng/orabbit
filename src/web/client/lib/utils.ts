import { LockerModel } from "@/actions/locker";
import { type ClassValue, clsx } from "clsx";
import _ from "lodash";
import toast from "react-hot-toast";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export async function standardToastableAction<T>(
  action: () => Promise<T>,
  loadingMessage: React.ReactElement,
  successMessage: React.ReactElement,
  errorMessage: React.ReactElement,
  successCallbacks: (() => void)[],
  errorCallbacks: (() => void)[]
) {
  const toastId = toast.loading(loadingMessage);
  try {
    await action();
    toast.success(successMessage, {
      id: toastId,
    });
    successCallbacks.forEach((callback) => callback());
  } catch (e) {
    toast.error(errorMessage, { id: toastId });
    errorCallbacks.forEach((callback) => callback());
  }
}

/**
 * Get the active locker from a list of lockers
 * @param lockers List of lockers
 * @returns The active locker or null if none is active
 */
export function GetActiveLocker(lockers: LockerModel[]) {
  if (!lockers) return null;

  const orderedEnableLocker = _.sortBy(
    lockers.filter((locker) => locker?.Enabled),
    (locker) => locker?.UpdatedAt,
    "desc"
  );
  if (orderedEnableLocker.length === 0) return null;
  return orderedEnableLocker[0];
}
