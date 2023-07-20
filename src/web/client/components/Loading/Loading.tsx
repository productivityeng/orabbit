import { Loader, Loader2 } from "lucide-react";
import React from "react";

function Loading() {
  return (
    <Loader2 className="animate-spin w-full h-full text-rabbit duration-1000" />
  );
}

export default Loading;
