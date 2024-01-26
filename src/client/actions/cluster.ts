"use server";
import { RabbitMqCluster } from "@/types";
import { FrontResponse, PaginatedResponse } from "./common/frontresponse";
import { CreateRabbitMqClusterRequestSchema } from "@/schemas/cluster-schemas";
import { boolean, z } from "zod";

export type FetchAllClustersResult = {
  result: RabbitMqCluster[];
  pageNumber: 1;
  pageSize: 100;
  totalItems: 3;
};
export async function fetchAllClusters() {
  //todo: build a better method for retrieve all clusters
  let result = await fetch(
    `${process.env
      .PRIVATE_INVENTORY_ENDPOINT!}/cluster?PageNumber=1&PageSize=100`,
    {
      method: "GET",
      cache: "no-store",
    }
  );
  let payloadResult = await result.json();
  let finalResult = payloadResult as RabbitMqCluster[];
  return finalResult;
}

export async function fetchCluster(clusterId: number) {
  let response = await fetch(
    `${process.env.PRIVATE_INVENTORY_ENDPOINT}/cluster/${clusterId}`,
    {
      cache: "no-store",
    }
  );

  switch (response.status) {
    case 200:
      let contentResponse = (await response.json()) as RabbitMqCluster;
      return contentResponse;
  }

  throw new Error("Fail to find cluster");
}

export async function createNewCluster(
  request: z.infer<typeof CreateRabbitMqClusterRequestSchema>
): Promise<FrontResponse<RabbitMqCluster | null>> {
  let response = await fetch(
    `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/cluster  `,
    {
      body: JSON.stringify(request),
      method: "POST",
    }
  );
  switch (response.status) {
    case 201: {
      let contentResponse = (await response.json()) as RabbitMqCluster;
      return { ErrorMessage: null, Result: contentResponse };
    }

    case 400: {
      let contentBadRequest = (await response.json()) as { error: string };
      return { ErrorMessage: contentBadRequest.error, Result: null };
    }

    case 406: {
      let contentInaceptable = (await response.json()) as {
        error: string;
        field: string;
      };
      return {
        ErrorMessage: `field ${contentInaceptable.field} with error => ${contentInaceptable.error}`,
        Result: null,
      };
    }

    case 500: {
      let contenctUnkow = await response.json();
      return { ErrorMessage: JSON.stringify(contenctUnkow), Result: null };
    }
    default:
      throw new Error("Erro desconhecido => " + JSON.stringify(response));
  }
}

export const DeleteCluster = async (clusterId: number): Promise<Boolean> => {
  let response = await fetch(
    `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/cluster/${clusterId}`,
    { method: "DELETE", cache: "no-store" }
  );
  if (response.status == 204) return true;

  console.log(await response.json());
  return false;
};
