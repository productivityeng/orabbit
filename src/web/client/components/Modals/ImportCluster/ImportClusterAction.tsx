"use client";
import React from "react";
import { Button } from "../../ui/button";
import { BoxesIcon } from "lucide-react";
import { Separator } from "../../ui/separator";
import { useTranslations } from "next-intl";
import Link from "next/link";

type Props = {
  ImportRoute: string;
};

/**
 * A nice card with an invitation for the user to bring a new cluster to the ecosystem. This component is container-based.
 * @param OnImportClick The function will be called when the Import button is clicked
 */
function ImportClusterAction({ ImportRoute }: Props) {
  const t = useTranslations();
  return (
    <div>
      <div className="flex flex-col group bg-slate-800 transition duration-300 hover:text-slate-200 text-white p-5 rounded-lg justify-center items-center ">
        <div className="flex">
          <BoxesIcon className="h-8 w-8 mr-4" />{" "}
          <h1 className="text-2xl">
            {t("Dashboard.ImportClusterButton.Title")}
          </h1>
        </div>
        <Separator className="my-1" />
        <p className="font-light text-justify">
          {t("Dashboard.ImportClusterButton.Body")}
        </p>

        <Link className="w-full" href={ImportRoute}>
          <Button className="bg-rabbit hover:bg-rabbit   active:scale-95  w-full mt-5 transition duration-300 text-lg">
            {t("Commons.Import")}
          </Button>
        </Link>
      </div>
    </div>
  );
}

export default ImportClusterAction;
