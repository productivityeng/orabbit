import React from "react";
import {
  AvatarFallback,
  AvatarImage,
  Avatar as ShadCnAvatar,
} from "../ui/avatar";

function Avatar() {
  return (
    <div>
      <ShadCnAvatar>
        <AvatarImage src="https://github.com/shadcn.png" alt="@shadcn" />
        <AvatarFallback>CN</AvatarFallback>
      </ShadCnAvatar>
    </div>
  );
}

export default Avatar;
