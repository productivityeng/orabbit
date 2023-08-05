"use client";
import React, { MouseEventHandler } from "react";
import { Button } from "../ui/button";
import { ArrowBigLeft, ArrowBigRight } from "lucide-react";
import { useTranslations } from "next-intl";

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
  const t = useTranslations();
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
        {t("Commons.Back")}
      </Button>
      <Button
        size="sm"
        disabled={isNextDisabled}
        onClick={OnNextClicked}
        className="rounded-sm bg-rabbit hover:bg-rabbit/90 active:scale-95 active:ring-rabbit active:ring-2 active:ring-offset-1 transition duration-200"
        variant="default"
      >
        {t("Commons.Next")}
        <ArrowBigRight />
      </Button>
    </div>
  );
}

export default NavigationBottom;
