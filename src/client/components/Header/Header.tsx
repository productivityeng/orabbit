import React from "react";
import Avatar from "./Avatar";
import { LanguageSelector } from "../cross/language-selector";

function Header() {
  return (
    <div className=" w-full h-full">
      <div className="flex flex-row justify-end px-10 items-center h-full">
        <div className="flex space-x-2">
          <LanguageSelector />
          <Avatar />
        </div>
      </div>
    </div>
  );
}

export default Header;
