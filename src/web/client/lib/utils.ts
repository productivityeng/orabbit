import { type ClassValue, clsx } from "clsx";
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
