// ImportClusterButton.test.js
import "@testing-library/jest-dom/extend-expect";

import React from "react";
import { render } from "@testing-library/react";
import ImportClusterAction from "./ImportClusterAction";

jest.mock("next-intl", () => ({
  useTranslations: jest.fn().mockReturnValue((value: string) => value),
}));

describe("ImportClusterButton", () => {
  it("renders ImportClusterButton component", () => {
    render(<ImportClusterAction ImportRoute="" />);
  });

  it("renders correct title", () => {
    const { getByText } = render(<ImportClusterAction ImportRoute="" />);

    const titleElement = getByText("Dashboard.ImportClusterButton.Title");
    expect(titleElement).toBeInTheDocument();
  });

  it("renders correct route for Import Route button", () => {
    const { getByText } = render(
      <ImportClusterAction ImportRoute="/importcluster" />
    );

    const titleElement = getByText("Dashboard.ImportClusterButton.Title");
    expect(titleElement).toBeInTheDocument();
  });
});
