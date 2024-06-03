import { render } from "@testing-library/react";
import "jest";
import { GenerateVirtualHostMq } from "@/__mocks__/models-generator";
import { VirtualHostsTable } from "./virtual-host-table";
import { RabbitMqVirtualHostColumnDef } from "./columns";
import { VirtualTableContext } from "./virtualhost-table-context";
import { act } from "react-dom/test-utils";

jest.mock("next/navigation", () => ({
  useRouter: () => ({
    query: { clusterId: "200" },
  }),
  useParams: () => ({ clusterId: "200" }),
}));

jest.mock("next-intl", () => ({
  useTranslations: () => jest.fn(),
}));

describe("should ", () => {
  const clusterNumber = 200;
  const clusterQtd = 10;

  it("render virtual-host-table", async () => {
    let vhost = GenerateVirtualHostMq(clusterQtd, clusterNumber);

    const { findAllByRole } = render(
      <VirtualTableContext.Provider value={{}}>
        <VirtualHostsTable
          columns={RabbitMqVirtualHostColumnDef}
          data={vhost}
        />
      </VirtualTableContext.Provider>
    );

    const rows = await findAllByRole("table-row");
    expect(rows).toHaveLength(clusterQtd);
  });

  it("call syncronize function when sync button is clicked", async () => {
    let vhost = GenerateVirtualHostMq(clusterQtd, clusterNumber);

    vhost[0].IsInDatabase = true;
    vhost[0].IsInCluster = false;
    vhost[0].Id = 1000;

    const syncFunctionMock = jest.fn();
    const { findByTestId } = render(
      <VirtualTableContext.Provider
        value={{
          OnSyncronizeVirtualHost: syncFunctionMock,
        }}
      >
        <VirtualHostsTable
          columns={RabbitMqVirtualHostColumnDef}
          data={vhost}
        />
      </VirtualTableContext.Provider>
    );

    let virtualHostRowCheckbox: HTMLElement = await findByTestId(
      `virtual-host-table-checkbox-${vhost[0].Id}`
    );

    act(() => {
      virtualHostRowCheckbox.click();
    });

    const syncVirtualHostButton = await findByTestId("sync-vhost-button");
    act(async () => {
      await syncVirtualHostButton.click();
    });
    expect(syncFunctionMock).toHaveBeenCalled();
  });

  it("call import function when import button is clicked", async () => {
    let vhost = GenerateVirtualHostMq(clusterQtd, clusterNumber);

    vhost[0].IsInDatabase = false;
    vhost[0].IsInCluster = true;
    vhost[0].Id = 1000;

    const importFunctionMock = jest.fn();
    const { findByTestId } = render(
      <VirtualTableContext.Provider
        value={{
          OnImportVirtualHostClick: importFunctionMock,
        }}
      >
        <VirtualHostsTable
          columns={RabbitMqVirtualHostColumnDef}
          data={[
            {
              ...vhost[0],
              IsInDatabase: false,
              IsInCluster: true,
              Id: 1000,
            },
          ]}
        />
      </VirtualTableContext.Provider>
    );

    const virtualHostRowCheckbox = await findByTestId(
      `virtual-host-table-checkbox-${1000}`
    );

    act(() => {
      virtualHostRowCheckbox.click();
    });

    const importButton = await findByTestId("import-vhost-button");
    act(async () => {
      await importButton.click();
    });
    expect(importFunctionMock).toHaveBeenCalled();
  });
});
