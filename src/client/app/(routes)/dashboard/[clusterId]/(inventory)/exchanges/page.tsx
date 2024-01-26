import { fetchExchangesFromCluster } from "@/actions/exchanges";
import Heading from "@/components/Heading/Heading";
import { Router } from "lucide-react";
import React from "react";
import { ExchangeTable } from "./components/exchange-table/exchange-table";
import { RabbitMqExchangeColumnDef } from "./components/exchange-table/columns";
import SimpleHeading from "@/components/Heading/SimpleHeading";

async function ExchangesPage({ params }: { params: { clusterId: number } }) {
  const exchanges = await fetchExchangesFromCluster(params.clusterId);
  return (
    <main>
      <SimpleHeading
        title={"Exchanges rastreadas"}
        description={"Todas as exchanges visiveis no cluster e na base"}
      />
      <ExchangeTable
        columns={RabbitMqExchangeColumnDef}
        data={exchanges.Result ?? []}
        searchKey="Name"
      />
    </main>
  );
}

export default ExchangesPage;
