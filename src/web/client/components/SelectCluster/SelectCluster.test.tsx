import React from "react";
import { render, fireEvent, screen, waitFor } from "@testing-library/react";
import { useRouter } from "next/navigation"; // Mock do useRouter
import { RabbitMqCluster } from "@/types";
import "@testing-library/jest-dom/extend-expect";
import { SelectCluster } from "./SelectCluster";
import { faker } from "@faker-js/faker";
import ResizeObserver from "resize-observer-polyfill";
global.ResizeObserver = ResizeObserver;

jest.mock("next/navigation", () => ({
  useRouter: jest.fn().mockReturnValue({
    push: jest.fn(),
  }),
}));

jest.mock("next-intl", () => ({
  useTranslations: jest.fn().mockReturnValue((value: string) => value),
}));

let mockClusters: RabbitMqCluster[] = [
  {
    ID: 1,
    CreatedAt: new Date("2023-07-19T00:00:00Z"),
    updatedAt: new Date("2023-07-19T12:00:00Z"),
    deletedAt: null,
    name: "Cluster 1",
    description: "Cluster 1 description",
    host: "host1.example.com",
    port: 5672,
    user: "user1",
    password: faker.internet.password(),
  },
  {
    ID: 2,
    CreatedAt: new Date("2023-07-20T00:00:00Z"),
    updatedAt: new Date("2023-07-20T12:00:00Z"),
    deletedAt: null,
    name: "Cluster 2",
    description: "Cluster 2 description",
    host: "host2.example.com",
    port: 5672,
    user: "user2",
    password: faker.internet.password(),
  },
  {
    ID: 3,
    CreatedAt: new Date("2023-07-21T00:00:00Z"),
    updatedAt: new Date("2023-07-21T12:00:00Z"),
    deletedAt: null,
    name: "Cluster 3",
    description: "Cluster 3 description",
    host: "host3.example.com",
    port: 5672,
    user: "user3",
    password: faker.internet.password(),
  },
];

describe("SelectCluster component", () => {
  it("should render the initial state correctly", () => {
    render(
      <SelectCluster
        SelectedCluster={undefined}
        Clusters={mockClusters}
        SetSelectedClusterId={jest.fn()}
      />
    );

    expect(screen.getByRole("combobox")).toBeInTheDocument();
    expect(screen.getByText("ClusterSelect...")).toBeInTheDocument();
    expect(screen.queryByText("Cluster 1")).not.toBeInTheDocument();
    expect(screen.queryByText("Cluster 2")).not.toBeInTheDocument();
    expect(screen.queryByText("Cluster 3")).not.toBeInTheDocument();
  });

  it("should display the cluster names when the button is clicked", () => {
    render(
      <SelectCluster
        SelectedCluster={undefined}
        Clusters={mockClusters}
        SetSelectedClusterId={jest.fn()}
      />
    );

    fireEvent.click(screen.getByRole("combobox"));

    expect(screen.getByText("Cluster 1")).toBeInTheDocument();
    expect(screen.getByText("Cluster 2")).toBeInTheDocument();
    expect(screen.getByText("Cluster 3")).toBeInTheDocument();
  });

  it("should select a cluster and navigate to the dashboard when clicked", async () => {
    const { push } = useRouter();
    render(
      <SelectCluster
        SelectedCluster={undefined}
        Clusters={mockClusters}
        SetSelectedClusterId={jest.fn()}
      />
    );

    fireEvent.click(screen.getByRole("combobox"));
    fireEvent.click(screen.getByText("Cluster 2"));
    await waitFor(() => {
      expect(push).toHaveBeenCalledWith("/dashboard/2");
    });
  });

  it("should close the popover when a cluster is selected", async () => {
    render(
      <SelectCluster
        SelectedCluster={undefined}
        Clusters={mockClusters}
        SetSelectedClusterId={jest.fn()}
      />
    );

    fireEvent.click(screen.getByRole("combobox"));
    fireEvent.click(screen.getByText("Cluster 1"));

    expect(screen.queryByText("Cluster 2")).not.toBeInTheDocument();
    expect(screen.queryByText("Cluster 3")).not.toBeInTheDocument();
    expect(screen.queryByText("Cluster 1")).not.toBeInTheDocument();
  });

  it("should render select cluster in combobox when is loaded selected", async () => {
    render(
      <SelectCluster
        SelectedCluster={mockClusters[2]}
        Clusters={mockClusters}
        SetSelectedClusterId={jest.fn()}
      />
    );

    await waitFor(() => {
      expect(screen.queryAllByText("Cluster 3").length).toBeGreaterThan(0);
    });
  });

  it("should call zustand to store the clusterId when cluster is selected", async () => {
    const SetSelectedClusterIdMock = jest.fn();

    render(
      <SelectCluster
        SelectedCluster={mockClusters[2]}
        Clusters={mockClusters}
        SetSelectedClusterId={SetSelectedClusterIdMock}
      />
    );

    fireEvent.click(screen.getByRole("combobox"));
    fireEvent.click(screen.getByText("Cluster 1"));

    await waitFor(() => {
      expect(SetSelectedClusterIdMock).toBeCalled();
    });
  });

  it("should redirect to home dashboard when select a cluster that is already selected", async () => {
    const { push } = useRouter();

    const SetSelectedClusterIdMock = jest.fn();

    render(
      <SelectCluster
        SelectedCluster={mockClusters[0]}
        Clusters={mockClusters}
        SetSelectedClusterId={SetSelectedClusterIdMock}
      />
    );

    fireEvent.click(screen.getByRole("combobox"));
    fireEvent.click(screen.getAllByText("Cluster 1")[1]);

    await waitFor(() => {
      expect(push).toHaveBeenCalledWith(`/dashboard`);
      expect(SetSelectedClusterIdMock).toHaveBeenCalledWith(undefined);
    });
  });
});
