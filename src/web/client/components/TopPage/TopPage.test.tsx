import { render } from "@testing-library/react";
import TopPage from "./TopPage";
import "@testing-library/jest-dom/extend-expect";
import { UsersIcon } from "lucide-react";

jest.mock("next-intl", () => ({
  useTranslations: jest.fn(() => (key: string) => key),
}));

describe("TopPage component", () => {
  it("renders without crashing", () => {
    render(
      <TopPage
        Icon={UsersIcon}
        IconColor="red"
        BgIconColor="blue"
        Titlei18Label="Title"
        Descriptioni18Label="Description"
      />
    );
  });

  it("displays the correct title and description", () => {
    const title = "Test Title";
    const description = "Test Description";
    const { getByText } = render(
      <TopPage
        Icon={UsersIcon}
        IconColor="red"
        BgIconColor="blue"
        Titlei18Label={title}
        Descriptioni18Label={description}
      />
    );
    expect(getByText(title)).toBeInTheDocument();
    expect(getByText(description)).toBeInTheDocument();
  });

  it("displays the provided icon", () => {
    const { container } = render(
      <TopPage
        Icon={UsersIcon}
        IconColor="red"
        BgIconColor="blue"
        Titlei18Label="Title"
        Descriptioni18Label="Description"
      />
    );
    const icon = container.querySelector("svg");
    expect(icon).toBeInTheDocument();
  });

  it("displays the translated title and description", () => {
    const titleKey = "title.key";
    const descriptionKey = "description.key";
    const { getByText } = render(
      <TopPage
        Icon={UsersIcon}
        IconColor="red"
        BgIconColor="blue"
        Titlei18Label={titleKey}
        Descriptioni18Label={descriptionKey}
      />
    );

    expect(getByText(titleKey)).toBeInTheDocument();
    expect(getByText(descriptionKey)).toBeInTheDocument();
  });
});
