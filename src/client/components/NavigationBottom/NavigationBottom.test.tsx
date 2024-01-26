import "@testing-library/jest-dom/extend-expect";

import React from "react";
import { render, fireEvent } from "@testing-library/react";
import NavigationBottom from "./NavigationBottom";
jest.mock("next-intl", () => ({
  useTranslations: jest.fn().mockReturnValue((value: string) => value),
}));

describe("NavigationBottom", () => {
  it("should disable Back button when isBackDisabled prop is true", () => {
    const onBackClickMock = jest.fn();
    const { getByText } = render(
      <NavigationBottom
        isBackDisabled={true}
        OnBackClicked={onBackClickMock}
        OnNextClicked={() => {}}
      />
    );

    const backButton = getByText("Commons.Back");
    expect(backButton).toBeDisabled();

    fireEvent.click(backButton);
    expect(onBackClickMock).not.toHaveBeenCalled();
  });

  it("should disable Next button when isNextDisabled prop is true", () => {
    const onNextClickMock = jest.fn();
    const { getByText } = render(
      <NavigationBottom
        isNextDisabled={true}
        OnBackClicked={() => {}}
        OnNextClicked={onNextClickMock}
      />
    );

    const nextButton = getByText("Commons.Next");
    expect(nextButton).toBeDisabled();

    fireEvent.click(nextButton);
    expect(onNextClickMock).not.toHaveBeenCalled();
  });

  it("should trigger OnBackClicked callback when Back button is clicked", () => {
    const onBackClickMock = jest.fn();
    const { getByText } = render(
      <NavigationBottom
        isBackDisabled={false}
        OnBackClicked={onBackClickMock}
        OnNextClicked={() => {}}
      />
    );

    const backButton = getByText("Commons.Back");
    fireEvent.click(backButton);

    expect(onBackClickMock).toHaveBeenCalled();
  });

  it("should trigger OnNextClicked callback when Next button is clicked", () => {
    const onNextClickMock = jest.fn();
    const { getByText } = render(
      <NavigationBottom
        isNextDisabled={false}
        OnBackClicked={() => {}}
        OnNextClicked={onNextClickMock}
      />
    );

    const nextButton = getByText("Commons.Next");
    fireEvent.click(nextButton);

    expect(onNextClickMock).toHaveBeenCalled();
  });
});
