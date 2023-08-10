"use client";
import React, { useState } from "react";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useForm } from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
} from "../../ui/form";
import { Input } from "../../ui/input";
import { Textarea } from "../../ui/textarea";
import { Button } from "../../ui/button";
import { useRouter } from "next/navigation";
import { useTranslations } from "next-intl";
import { Frown, Loader2 } from "lucide-react";
import { toast } from "react-hot-toast";
import { cn } from "@/lib/utils";
import Modal from "../../ui/modal";
import { useImportCluster } from "@/hooks/import-cluster";
import { createNewCluster } from "@/services/cluster";
import { CreateRabbitMqClusterRequestSchema } from "@/schemas/CreateRabbitMqClusterRequestSchema";

function ImportClusterForm() {
  const t = useTranslations();
  const router = useRouter();
  const [creationError, setCreationError] = useState<string>();
  const { closeModal, isModalOpen } = useImportCluster();

  const form = useForm<z.infer<typeof CreateRabbitMqClusterRequestSchema>>({
    resolver: zodResolver(CreateRabbitMqClusterRequestSchema),
    defaultValues: {
      port: 15672,
      description: "",
      host: "",
      name: "",
      password: "",
      user: "",
    },
  });
  const submit = async (
    values: z.infer<typeof CreateRabbitMqClusterRequestSchema>
  ) => {
    try {
      setCreationError(undefined);
      const createCluster = async () => {
        let creaedCluster = await createNewCluster(values);
        if (creaedCluster.Result) {
          router.push(`/dashboard/${creaedCluster.Result.Id}`);
          closeModal();
          router.refresh();
        } else {
          console.error(creaedCluster.ErrorMessage);
          setCreationError(creaedCluster.ErrorMessage!);
          return Promise.reject(creaedCluster);
        }
      };

      await toast.promise(createCluster(), {
        loading: <b>{t("ImportClusterForm.Creating")}</b>,
        success: "Cluster created!",
        error: (
          <b className="flex flex-row gap-x-3">
            {t("ImportClusterForm.FailToCreate")}{" "}
            <Frown className="fill-yellow-500" />
          </b>
        ),
      });
    } catch (error) {
      console.error(error);
    }
  };
  return (
    <Modal
      title={t("Sidebar.ImportCluster")}
      description="Bring a new cluster to orabbit umbrella"
      isOpen={isModalOpen}
      onClose={closeModal}
    >
      <Form {...form}>
        <form
          role="form"
          onSubmit={form.handleSubmit(submit)}
          className="border-neutral-50 border-2  rounded-sm  p-5 grid"
        >
          <div className=" w-full">
            <p
              className={cn(
                "bg-red-100/50 text-center  rounded-sm text-red-500",
                {
                  "p-2": creationError,
                }
              )}
            >
              {creationError}
            </p>

            <FormField
              control={form.control}
              name="name"
              render={({ field, formState: { errors, isDirty } }) => (
                <FormItem>
                  <FormLabel>{t("Commons.Name")}</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="InvestBank compliance division"
                      disabled={form.formState.isSubmitting}
                    />
                  </FormControl>
                  <p className="text-red-500 font-light text-xs">
                    {errors.name?.message}
                  </p>
                </FormItem>
              )}
            />
          </div>

          <FormField
            control={form.control}
            name="description"
            render={({ field, formState: { errors } }) => (
              <FormItem>
                <FormLabel>{t("Commons.Description")}</FormLabel>
                <FormControl>
                  <Textarea
                    {...field}
                    placeholder="Manage all events of transactional banking with another areas of the bank"
                    disabled={form.formState.isSubmitting}
                    className=" w-full bg-white"
                  />
                </FormControl>
                <p className="text-red-500 font-light text-xs">
                  {errors.description?.message}
                </p>
              </FormItem>
            )}
          />
          <div className="w-full grid grid-cols-12 ">
            <div className="col-span-12 ">
              <FormField
                control={form.control}
                name="host"
                render={({ field, formState: { errors } }) => (
                  <FormItem>
                    <FormLabel>{t("Commons.Host")}</FormLabel>
                    <FormControl>
                      <Input
                        placeholder="rabbit.com"
                        disabled={form.formState.isSubmitting}
                        {...field}
                      />
                    </FormControl>
                    <p className="text-red-500 font-light text-xs">
                      {errors.host?.message}
                    </p>
                  </FormItem>
                )}
              />
            </div>

            <div className="col-span-12 ">
              <FormField
                control={form.control}
                name="port"
                render={({ field, formState: { errors } }) => (
                  <FormItem className="grid-cols-2">
                    <FormLabel>{t("Commons.Port")}</FormLabel>
                    <FormControl>
                      <Input
                        type="number"
                        disabled={form.formState.isSubmitting}
                        {...field}
                        onChange={(value) =>
                          field.onChange(+value.target.value)
                        }
                      />
                    </FormControl>
                    <p className="text-red-500 font-light text-xs">
                      {errors.port?.message}
                    </p>
                  </FormItem>
                )}
              />
            </div>
          </div>

          <div className="relative w-full lg:space-x-10 flex lg:flex-row flex-col justify-between">
            <FormField
              control={form.control}
              name="user"
              render={({ field, formState: { errors } }) => (
                <FormItem className="w-full">
                  <FormLabel>{t("Commons.User")}</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      placeholder="username"
                      disabled={form.formState.isSubmitting}
                    />
                  </FormControl>{" "}
                  <p className="text-red-500 font-light text-xs">
                    {errors.user?.message}
                  </p>
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field, formState: { errors } }) => (
                <FormItem className="w-full">
                  <FormLabel>{t("Commons.Password")}</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      disabled={form.formState.isSubmitting}
                      placeholder="password"
                    />
                  </FormControl>
                  <p className="text-red-500 font-light text-xs">
                    {errors.password?.message}
                  </p>
                </FormItem>
              )}
            />
          </div>
          <div className=" w-full py-5 flex justify-end space-x-5">
            <Button
              variant="ghost"
              size="sm"
              disabled={form.formState.isSubmitting}
              onClick={() => closeModal()}
              className="rounded-sm  active:scale-95 active:ring-slate-900 active:ring-2 active:ring-offset-1 transition duration-200"
            >
              {t("Commons.Cancel")}
            </Button>{" "}
            <Button
              type="submit"
              size="sm"
              disabled={!form.formState.isValid || form.formState.isSubmitting}
              className="rounded-sm px-5 space-x-2  bg-rabbit hover:bg-rabbit/90 active:scale-95 active:ring-rabbit active:ring-2 active:ring-offset-1 transition duration-200 "
            >
              {form.formState.isSubmitting ? (
                <Loader2 className="animate-spin" />
              ) : (
                <>{t("Commons.Create")}</>
              )}
            </Button>
          </div>
        </form>
      </Form>
    </Modal>
  );
}

export default ImportClusterForm;
