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
import { zodResolver } from "@hookform/resolvers/zod";
import { Frown, Import, Loader2, Plus } from "lucide-react";
import { useTranslations } from "next-intl";
import { useParams, useRouter } from "next/navigation";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "react-hot-toast";
import { z } from "zod";
import { CreateRabbitMqQeueueRequestSchema } from "@/schemas/queue-schemas";
import { createQueue } from "@/actions/queue";
import { RabbitMqQueue } from "@/models/queues";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

interface QueueFormProps {
  initialData: RabbitMqQueue | null;
}

function QueueForm({ initialData }: QueueFormProps) {
  const params = useParams() as unknown as { clusterId: number };
  const router = useRouter();
  const [creationError, setCreationError] = useState<string>();
  const t = useTranslations();

  const queueTypes = ["classic", "quorum"];

  const form = useForm<z.infer<typeof CreateRabbitMqQeueueRequestSchema>>({
    resolver: zodResolver(CreateRabbitMqQeueueRequestSchema),
    defaultValues: {
      ClusterId: parseInt(params.clusterId.toString()),
      QueueName: "",
      Create: false,
    },
  });

  const onSubmit = async (
    data: z.infer<typeof CreateRabbitMqQeueueRequestSchema>
  ) => {
    setCreationError("");
    let toastId = toast.loading("Importing a queue from cluster");
    try {
      const response = await createQueue(params.clusterId, data);

      if (response.Result) {
        router.push(
          `/dashboard/${params.clusterId}/queuesandstreams/${response.Result.ID}`
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
        className="flex flex-col w-full justify-center items-center"
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
              name="Create"
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
        <div className="w-1/2 space-y-5">
          <FormField
            control={form.control}
            name="QueueName"
            render={({ field }) => (
              <FormItem>
                <FormLabel> {t("ImportClusterForm.Queue.Label")}</FormLabel>
                <FormControl>
                  <Input
                    className="col-span-3"
                    {...field}
                    type="text"
                    placeholder={t("ImportClusterForm.Queue.Placeholder")}
                  />
                </FormControl>
                <FormDescription>
                  {t("ImportClusterForm.Queue.Description")}
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />

          <FormField
            name="Type"
            control={form.control}
            render={({ field }) => (
              <FormItem className="col-span-2 md:col-span-1">
                <FormLabel> {t("ImportClusterForm.Type.Label")}</FormLabel>
                <Select
                  disabled={form.formState.isLoading}
                  onValueChange={field.onChange}
                  value={field.value}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger className="bg-background">
                      <SelectValue
                        defaultValue={field.value}
                        placeholder={t("ImportClusterForm.Type.Placeholder")}
                      />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    {queueTypes.map((queueType) => (
                      <SelectItem key={queueType} value={queueType}>
                        {queueType}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                <FormDescription>
                  {t("ImportClusterForm.Type.Description")}
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
        </div>

        <div className="flex justify-end w-1/2 space-x-4 py-2">
          <Button
            variant="outline"
            type="button"
            onClick={() => router.push("./")}
          >
            Cancelar
          </Button>
          <Button className="" disabled={!form.formState.isValid}>
            {form.formState.isSubmitting && (
              <Loader2 className="w-4 h-4 mr-2 animate-spin" />
            )}
            {action == "Criar" && <Plus className="mr-2" />}
            {action == "Importar" && <Import className="mr-2" />}
            {action}
          </Button>
        </div>
      </form>
    </Form>
  );
}

export default QueueForm;
