import React from "react";
import Avatar from "./Avatar";

function Header() {
  return (
    <div className=" w-full h-full">
      <div className="flex flex-row justify-end px-10 items-center h-full">
        <div>
          <Avatar />
        </div>
      </div>
    </div>
  );
}

export default Header;
