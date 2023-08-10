import "@testing-library/jest-dom/extend-expect";

import {
  render,
  screen,
  fireEvent,
  waitFor,
  FireFunction,
  FireObject,
} from "@testing-library/react";
import { useForm, FormProvider } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import ImportClusterForm from "./ImportClusterModal";
import { CreateRabbitMqClusterRequest } from "@/models/cluster";
import React from "react";
import { act } from "react-dom/test-utils";
import { RabbitMqCluster } from "../../../types";
import { FrontResponse } from "@/services/common/frontresponse";
import { faker } from "@faker-js/faker";
const refresh = jest.fn();
const push = jest.fn();

jest.mock("next-intl", () => ({
  useTranslations: jest.fn().mockReturnValue((value: string) => value),
}));

jest.mock("next/navigation", () => ({
  useRouter: () => ({
    push: push,
    refresh: refresh,
  }),
}));

describe("ImportClusterForm", () => {
  const mockedValues: CreateRabbitMqClusterRequest = {
    name: faker.company.name(),
    description: faker.commerce.productDescription(),
    host: "testhost",
    port: 15672,
    user: "testuser",
    password: faker.internet.password(),
  };

  let mockResult: FrontResponse<RabbitMqCluster | null> = {
    Result: {
      Id: 99,
      description: mockedValues.description,
      host: mockedValues.description,
      port: mockedValues.port,
      user: mockedValues.user,
      password: mockedValues.password,
      createdAt: new Date(),
      updatedAt: new Date(),
      deletedAt: null,
      name: mockedValues.name,
    },
    ErrorMessage: null,
  };

  const formSchema = z.object({
    name: z.string().min(1),
    description: z.string().min(1),
    host: z.string().min(1),
    port: z.number().int().positive(),
    user: z.string().min(1),
    password: z.string().min(8),
  });

  const populateForm = (fireEvent: FireFunction & FireObject) => {
    fireEvent.change(screen.getByLabelText("Commons.Name"), {
      target: { value: mockedValues.name },
    });
    fireEvent.change(screen.getByLabelText("Commons.Description"), {
      target: { value: mockedValues.description },
    });
    fireEvent.change(screen.getByLabelText("Commons.Host"), {
      target: { value: mockedValues.host },
    });
    fireEvent.change(screen.getByLabelText("Commons.Port"), {
      target: { value: mockedValues.port },
    });
    fireEvent.change(screen.getByLabelText("Commons.User"), {
      target: { value: mockedValues.user },
    });
    fireEvent.change(screen.getByLabelText("Commons.Password"), {
      target: { value: mockedValues.password },
    });
  };

  const onSubmit = jest.fn().mockReturnValue(mockResult);
  const onCancel = jest.fn();

  const WrapperComponent = ({ children }: any) => {
    const form = useForm<CreateRabbitMqClusterRequest>({
      resolver: zodResolver(formSchema),
    });

    return (
      <FormProvider {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)}>{children}</form>
      </FormProvider>
    );
  };

  it("renders form fields", () => {
    render(
      <ImportClusterForm
        OnCreateClicked={onSubmit}
        OnCancelClicked={onCancel}
      />,
      {
        wrapper: WrapperComponent,
      }
    );

    expect(screen.getByLabelText("Commons.Name")).toBeInTheDocument();
    expect(screen.getByLabelText("Commons.Description")).toBeInTheDocument();
    expect(screen.getByLabelText("Commons.Host")).toBeInTheDocument();
    expect(screen.getByLabelText("Commons.Port")).toBeInTheDocument();
    expect(screen.getByLabelText("Commons.User")).toBeInTheDocument();
  });

  it("shows error messages for empty form fields on submit", async () => {
    render(
      <ImportClusterForm
        OnCreateClicked={onSubmit}
        OnCancelClicked={onCancel}
      />,
      {
        wrapper: WrapperComponent,
      }
    );
    act(() => {
      fireEvent.submit(screen.getByRole("form"));
    });

    await waitFor(() => {
      expect(
        screen.getByText("Validations.ImportClustterForm.Description")
      ).toBeInTheDocument();
      expect(
        screen.getByText("Validations.ImportClustterForm.Host")
      ).toBeInTheDocument();
      expect(
        screen.getByText("Validations.ImportClustterForm.Name")
      ).toBeInTheDocument();
      expect(
        screen.getByText("Validations.ImportClustterForm.Password")
      ).toBeInTheDocument();

      expect(
        screen.getByText("Validations.ImportClustterForm.User")
      ).toBeInTheDocument();
    });

    expect(onSubmit).not.toHaveBeenCalled();
  });

  it("not redirect when fail to create a cluster", async () => {
    const onSubmit = jest
      .fn()
      .mockReturnValue({ ErrorMessage: "Error", Result: null });

    jest.mock("next/navigation", () => ({
      useRouter: () => ({
        push: jest.fn(),
        refresh: refresh,
      }),
    }));
    render(
      <ImportClusterForm
        OnCreateClicked={onSubmit}
        OnCancelClicked={onCancel}
      />,
      {
        wrapper: WrapperComponent,
      }
    );

    act(() => {
      populateForm(fireEvent);
    });
    let createButton: HTMLElement;
    await waitFor(() => {
      createButton = screen.getByText("Commons.Create");
      expect(createButton).not.toBeNull();
      expect(createButton).not.toBeDisabled();
    });

    fireEvent.click(createButton!);
    await waitFor(() => {
      expect(onSubmit).toBeCalledWith(mockedValues);
      expect(refresh).not.toHaveBeenCalled();
      expect(push).not.toHaveBeenCalled();
    });
  });

  it("", async () => {
    jest.mock("next/navigation", () => ({
      useRouter: () => ({
        push: jest.fn(),
        refresh: refresh,
      }),
    }));
    render(
      <ImportClusterForm
        OnCreateClicked={onSubmit}
        OnCancelClicked={onCancel}
      />,
      {
        wrapper: WrapperComponent,
      }
    );

    act(() => {
      populateForm(fireEvent);
    });
    let createButton: HTMLElement;
    await waitFor(() => {
      createButton = screen.getByText("Commons.Create");
      expect(createButton).not.toBeNull();
      expect(createButton).not.toBeDisabled();
    });

    fireEvent.click(createButton!);
    await waitFor(() => {
      expect(onSubmit).toBeCalledWith(mockedValues);
      expect(refresh).toHaveBeenCalledTimes(1);
      expect(push).toHaveBeenCalledWith(`/dashboard/${mockResult.Result?.Id}`);
    });
  });

  it("calls the OnCancelClicked function when Cancel button is clicked", async () => {
    render(
      <ImportClusterForm
        OnCreateClicked={onSubmit}
        OnCancelClicked={onCancel}
      />,
      {
        wrapper: WrapperComponent,
      }
    );
    act(() => {
      fireEvent.click(screen.getByText("Commons.Cancel"));
    });
    await waitFor(() => {
      expect(onCancel).toHaveBeenCalled();
    });
  });
});
