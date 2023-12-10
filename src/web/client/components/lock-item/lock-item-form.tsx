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
import { useTranslations } from "next-intl";
import { extend } from "lodash";

const FormSchema = z.object({
  reason: z.string().min(10, {
    message: "The reason must be at least 10 characters",
  }),
});

interface LockItemFormProps extends React.HTMLAttributes<HTMLFormElement> {
  onFormSubmit: (data: z.infer<typeof FormSchema>) => void;
}
function LockItemForm({ onFormSubmit, ...props }: LockItemFormProps) {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {},
  });
  const t = useTranslations();

  return (
    <Form {...props} {...form} data-testid="lock-item-form">
      <form onSubmit={form.handleSubmit(onFormSubmit)}>
        <FormField
          control={form.control}
          
          name="reason"
          render={({ field }) => (
            <FormItem>
              <FormLabel role="heading">{t("label-reason")}</FormLabel>
              <FormControl>
                <Textarea
                  data-testid="reason-textarea"
                  className="resize-none"
                  {...field}
                />
              </FormControl>
              <FormDescription>{t("label-reason-explain")}</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button data-testid="submit-button" size="sm" className="float-right" type="submit">
          Submit
        </Button>
      </form>
    </Form>
  );
}

export default LockItemForm;
