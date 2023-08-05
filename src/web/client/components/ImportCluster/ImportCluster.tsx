"use client";
import { BoxesIcon } from "lucide-react";
import React from "react";
import { Separator } from "../ui/separator";
import { useTranslations } from "next-intl";
import ImportClusterForm from "./ImportClusterForm";
import { useRouter } from "next/navigation";
import { createNewCluster } from "@/services/cluster";

function ImportCluster() {
  const t = useTranslations();
  const router = useRouter();

  return (
    <div className="flex flex-col h-full w-full group    transition duration-300  text-slate-900 p-5 rounded-lg ">
      <div className="flex-grow">
        <div className="flex">
          <BoxesIcon className="h-8 w-8 mr-4" />{" "}
          <h1 className="text-2xl">{t("Dashboard.ImportCluster.Title")}</h1>
        </div>
        <Separator className="my-1" />
        <p className="font-light text-justify text-sm">
          {t("Dashboard.ImportCluster.Body")}
        </p>
        <div className="mt-5">
          {" "}
          <ImportClusterForm
            OnCancelClicked={() => router.push("/dashboard")}
            OnCreateClicked={createNewCluster}
          />
        </div>
      </div>

      {/* <div className="flex justify-end   w-full ">
        <NavigationBottom
          isBackDisabled={StepOrder[step].Previous == step}
          isNextDisabled={StepOrder[step].Next == step}
          OnBackClicked={() => setStep(StepOrder[step].Previous)}
          OnNextClicked={() => setStep(StepOrder[step].Next)}
        />
      </div> */}
    </div>
  );
}

export default ImportCluster;
