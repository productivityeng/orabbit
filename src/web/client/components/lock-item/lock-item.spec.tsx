import React from "react";
import { render, fireEvent, waitFor, act } from "@testing-library/react";
import LockItem from "./lock-item";
import { IntlProvider } from "next-intl";
import "@testing-library/jest-dom/extend-expect";
import toast from "react-hot-toast";
import { useTranslations } from "next-intl";

jest.mock("next-intl", () => ({
  ...jest.requireActual("next-intl"),
  useTranslations: () => (t: string) => t,
}));

describe("LockItem", () => {
  it("should display the dialog when clicking on the lock icon", () => {
    const { getByTestId, getByText } = render(
      <IntlProvider locale="en">
        <LockItem isLocked={false} lockType="User" artifactName="John Doe" />
      </IntlProvider>
    );

    const lockIcon = getByTestId("lock-unlock-button");
    act(() => {
      fireEvent.click(lockIcon);
    });

    const dialogTitle = getByText(/John Doe/i);
    expect(dialogTitle).toBeInTheDocument();
  });

  it("should call the submit function when a valid form is submitted", async () => {
    jest.mock("react-hot-toast");
    toast.success = jest.fn();
    const { getByTestId } = render(
      <IntlProvider locale="en">
        <LockItem isLocked={false} lockType="User" artifactName="John Doe" />
      </IntlProvider>
    );

    const lockIcon = getByTestId("lock-unlock-button");
    act(() => {
      fireEvent.click(lockIcon);
    });

    const textarea = getByTestId("reason-textarea");
    act(() => {
      fireEvent.change(textarea, {
        target: { value: "Texto com mais de 10 caracteres" },
      });
    });

    const submitButton = getByTestId("submit-button");
    act(() => {
      fireEvent.click(submitButton);
    });
    await waitFor(() => {
      expect(toast.success).toHaveBeenCalledTimes(1);
    });
  });

  it("should not call the submit function when a invalid form is submitted", async () => {
    jest.mock("react-hot-toast");
    toast.success = jest.fn();
    const { getByTestId } = render(
      <IntlProvider locale="en">
        <LockItem isLocked={false} lockType="User" artifactName="John Doe" />
      </IntlProvider>
    );

    act(() => {
      const lockIcon = getByTestId("lock-unlock-button");
      fireEvent.click(lockIcon);
    });
    act(() => {
      const textarea = getByTestId("reason-textarea");

      fireEvent.change(textarea, {
        target: { value: "Texto " },
      });
    });

    act(() => {
      const submitButton = getByTestId("submit-button");
      fireEvent.click(submitButton);
    });

    await waitFor(() => {
      expect(toast.success).toHaveBeenCalledTimes(0);
    });
  });
});
