"use client";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import { cn } from "@/lib/utils";
import { CreateRabbitmqUserSchema } from "@/schemas/user-schemas";
import { createUser } from "@/actions/users";
import { zodResolver } from "@hookform/resolvers/zod";
import { Frown, Loader2 } from "lucide-react";
import { useTranslations } from "next-intl";
import { useParams, useRouter } from "next/navigation";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "react-hot-toast";
import { z } from "zod";
import { RabbitMqUser } from "@/models/users";

interface UserFormProps {
  initialData: RabbitMqUser | null;
}

function UserForm({ initialData }: UserFormProps) {
  const params = useParams() as unknown as { clusterId: number };
  const router = useRouter();
  const [creationError, setCreationError] = useState<string>();
  const t = useTranslations();

  const form = useForm<z.infer<typeof CreateRabbitmqUserSchema>>({
    resolver: zodResolver(CreateRabbitmqUserSchema),
    defaultValues: {
      clusterId: parseInt(params.clusterId.toString()),
      password: "",
      username: initialData?.Username,
      create: true,
    },
  });

  const onSubmit = async (data: z.infer<typeof CreateRabbitmqUserSchema>) => {
    let toastId = toast.loading("Criando usuario");
    try {
      const response = await createUser(params.clusterId, data);
      if (response.Result) {
        toast.success("Cluster created!", { id: toastId });
        router.push(`/dashboard/${params.clusterId}/users/`);
        router.refresh();
      } else {
        console.error(response.ErrorMessage);
        setCreationError(response.ErrorMessage!);
        toast.error(
          <>
            {t("ImportClusterForm.FailToCreate")}{" "}
            <Frown className="fill-yellow-500" />
          </>,
          {
            id: toastId,
          }
        );
      }
    } catch (error) {
      console.error(error);
      toast.error(
        <>
          {t("ImportClusterForm.FailToCreate")}{" "}
          <Frown className="fill-yellow-500" />
        </>,
        {
          id: toastId,
        }
      );
    }
  };

  const action = initialData
    ? "Salvar"
    : form.getValues().create
    ? "Criar"
    : "Importar";

  return (
    <Form {...form}>
      <form
        role="form"
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col w-1/2"
      >
        <p
          className={cn(
            "bg-red-100/50 text-center text-sm  rounded-sm text-red-500 w-1/2 my-2",
            {
              "p-2": creationError,
            }
          )}
        >
          {t(creationError)}
        </p>

        {!initialData && (
          <div className=" py-2">
            <FormField
              control={form.control}
              name="create"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center space-x-4 justify-between rounded-lg border p-4">
                  <div className="space-y-0.5">
                    <FormLabel className="text-base">Criar</FormLabel>
                    <FormDescription className="text-justify">
                      Ao criar um novo usuario ele sera adicionado ao cluster de
                      rabbitMQ. Caso o usuario ja existe o mesmo não será
                      sobrescrito, use a opcao de importar o usuario caso ele ja
                      exista, não é necessário fornecer a senha para esta opção.
                    </FormDescription>
                  </div>
                  <FormControl>
                    <Switch
                      disabled
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
          </div>
        )}
        <div>
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input
                    className="col-span-3"
                    {...field}
                    type="text"
                    placeholder="username"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        {(form.getValues().create || initialData) && (
          <div className="py-2 transition-all duration-200">
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input
                      className="col-span-3"
                      {...field}
                      type="password"
                      placeholder="password"
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        )}
        <div className="flex justify-end  space-x-4 py-2">
          <Button
            variant="outline"
            type="button"
            size="sm"
            onClick={() => router.push("./")}
          >
            Cancelar
          </Button>
          <Button disabled={!form.formState.isValid} size="sm">
            {form.formState.isSubmitting && (
              <Loader2 className="w-4 h-4 mr-2 animate-spin" />
            )}
            {action}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export default UserForm;
