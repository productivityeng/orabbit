"use client";
import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "../ui/form";
import { Textarea } from "../ui/textarea";
import { Button } from "../ui/button";
import { extend } from "lodash";
import { LockItemFormSchema } from "@/schemas/locker-item-schemas";

interface LockItemFormProps extends React.HTMLAttributes<HTMLFormElement> {
  onFormSubmit: (data: z.infer<typeof LockItemFormSchema>) => Promise<void>;
}
function LockItemForm({ onFormSubmit, ...props }: LockItemFormProps) {
  const form = useForm<z.infer<typeof LockItemFormSchema>>({
    resolver: zodResolver(LockItemFormSchema),
    defaultValues: {},
  });
  return (
    <Form {...props} {...form} data-testid="lock-item-form">
      <form onSubmit={form.handleSubmit(onFormSubmit)}>
        <FormField
          control={form.control}
          name="reason"
          render={({ field }) => (
            <FormItem>
              <FormLabel role="heading">Motivo</FormLabel>
              <FormControl>
                <Textarea
                  data-testid="reason-textarea"
                  className="resize-none"
                  {...field}
                />
              </FormControl>
              <FormDescription>Explique o motivo do bloqueio</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button
          data-testid="submit-button"
          size="sm"
          className="float-right"
          type="submit"
        >
          Enviar
        </Button>
      </form>
    </Form>
  );
}

export default LockItemForm;
