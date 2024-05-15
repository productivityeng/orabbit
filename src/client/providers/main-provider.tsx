"use client";

import { queryClient } from "@/lib/query";
import { QueryClientProvider } from "@tanstack/react-query";

export const MainProvider = ({ children }: { children: React.ReactNode }) => { 
    return <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>;
};