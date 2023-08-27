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
import { RabbitMqQueue, RabbitMqUser } from "@/types";
import { zodResolver } from "@hookform/resolvers/zod";
import { Frown, Loader2 } from "lucide-react";
import { useTranslations } from "next-intl";
import { useParams, useRouter } from "next/navigation";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "react-hot-toast";
import { z } from "zod";
import { CreateRabbitMqQeueueRequestSchema } from "@/schemas/queue-schemas";

interface UserFormProps {
  initialData: RabbitMqQueue | null;
}

function UserForm({ initialData }: UserFormProps) {
  const params = useParams() as unknown as { clusterId: number };
  const router = useRouter();
  const [creationError, setCreationError] = useState<string>();
  const t = useTranslations();

  const form = useForm<z.infer<typeof CreateRabbitMqQeueueRequestSchema>>({
    resolver: zodResolver(CreateRabbitMqQeueueRequestSchema),
    defaultValues: {
      ClusterId: parseInt(params.clusterId.toString()),
      QueueName: "",
    },
  });

  const onSubmit = async (
    data: z.infer<typeof CreateRabbitMqQeueueRequestSchema>
  ) => {
    let toastId = toast.loading("Importing a queue from cluster");
    try {
      const response = await createUser(params.clusterId, data);

      if (response.Result) {
        router.push(
          `/dashboard/${params.clusterId}/queuesandstreams/${response.Result.Id}`
        );
        toast.success("Queue created!", { id: toastId });
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
    : form.getValues().Create
    ? "Criar"
    : "Importar";

  return (
    <Form {...form}>
      <form
        role="form"
        onSubmit={form.handleSubmit(onSubmit)}
        className="flex flex-col"
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
          <div className="w-1/2 py-2">
            <FormField
              control={form.control}
              name="Import"
              render={({ field }) => (
                <FormItem className="flex flex-row items-center space-x-4 justify-between rounded-lg border p-4">
                  <div className="space-y-0.5">
                    <FormLabel className="text-base">Criar</FormLabel>
                    <FormDescription className="text-justify">
                      Ao criar uma nova fila ela sera adicionada ao cluster de
                      rabbitMQ. Caso a fila ja exista ela não será sobrescrita,
                      use a opcao de importar fila caso ela ja exista.
                    </FormDescription>
                  </div>
                  <FormControl>
                    <Switch
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                  </FormControl>
                </FormItem>
              )}
            />
          </div>
        )}
        <div className="w-1/2">
          <FormField
            control={form.control}
            name="QueueName"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Queue Name</FormLabel>
                <FormControl>
                  <Input
                    className="col-span-3"
                    {...field}
                    type="text"
                    placeholder="queue name"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="flex justify-end w-1/2 space-x-4 py-2">
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
