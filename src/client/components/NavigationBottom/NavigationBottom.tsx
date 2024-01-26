"use client";
import React, { MouseEventHandler } from "react";
import { Button } from "../ui/button";
import { ArrowBigLeft, ArrowBigRight } from "lucide-react";

type Props = {
  isBackDisabled?: boolean;
  isNextDisabled?: boolean;
  OnBackClicked: MouseEventHandler;
  OnNextClicked: MouseEventHandler;
};
function NavigationBottom({
  isBackDisabled,
  isNextDisabled,
  OnBackClicked,
  OnNextClicked,
}: Props) {
  return (
    <div className="flex space-x-5">
      <Button
        variant="ghost"
        size="sm"
        disabled={isBackDisabled}
        onClick={OnBackClicked}
        className="rounded-sm  active:scale-95 active:ring-slate-900 active:ring-2 active:ring-offset-1 transition duration-200"
      >
        <ArrowBigLeft />
        Voltar
      </Button>
      <Button
        size="sm"
        disabled={isNextDisabled}
        onClick={OnNextClicked}
        className="rounded-sm bg-rabbit hover:bg-rabbit/90 active:scale-95 active:ring-rabbit active:ring-2 active:ring-offset-1 transition duration-200"
        variant="default"
      >
        Proximo
        <ArrowBigRight />
      </Button>
    </div>
  );
}

export default NavigationBottom;
