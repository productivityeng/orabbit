"use server";
import { RabbitMqQueue } from "@/models/queues";
import { FrontResponse } from "./common/frontresponse";
import { RabbitMqExchange } from "@/models/exchange";

export async function fetchExchangesFromCluster(
  clusterId: number
): Promise<FrontResponse<RabbitMqExchange[] | null>> {
  const fetchQueueFromClusterEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/exchange`;

  let response = await fetch(fetchQueueFromClusterEndpoint, {
    method: "GET",
    cache: "no-store",
  });

  switch (response.status) {
    case 200:
      return {
        ErrorMessage: null,
        Result: (await response.json()) as RabbitMqExchange[],
      };
    default:
      return {
        ErrorMessage: ((await response.json()) as { error: string }).error,
        Result: null,
      };
  }
}

/**
 * Import a exchange that exists in rabbitmq cluster and register it on ostern
 * @param clusterId
 * @param exchangeName
 */
export async function importExchangeFromClusterAction(
  clusterId: number,
  exchangeName: string
) {
  console.info(`Importing exchange ${exchangeName} from cluster ${clusterId}`);

  const importExchangeFromClusterEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/exchange/import`;

  console.info(`Sending request to ${importExchangeFromClusterEndpoint}`);
  let result = await fetch(importExchangeFromClusterEndpoint, {
    method: "POST",
    cache: "no-store",
    body: JSON.stringify({
      Name: exchangeName,
    }),
  });

  console.info(
    `Received response ${result.status} from ${importExchangeFromClusterEndpoint} `
  );

  switch (result.status) {
    case 200:
    case 201:
      return {
        ErrorMessage: null,
        Result: await result.json(),
      };
    default:
      return {
        ErrorMessage: ((await result.json()) as { error: string }).error,
        Result: null,
      };
  }
}

/**
 * Removes a exchange from a specified RabbitMQ Cluster based on ClusterID and ExchangeId.
 * @param ClusterId
 * @param ExchangeId
 * @returns
 */
export async function removeExchangeFromClusterAction(
  ClusterId: number,
  ExchangeId: number
): Promise<FrontResponse<boolean>> {
  console.log(`Removing exchange ${ExchangeId} from cluster ${ClusterId}`);

  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${ClusterId}/exchange/${ExchangeId}`;

  console.log(`Sending request to ${createUserEndpoint}`);
  let response = await fetch(createUserEndpoint, {
    method: "DELETE",
    cache: "no-store",
  });

  console.log(
    `Received response ${response.status} from ${createUserEndpoint} `
  );

  switch (response.status) {
    case 201:
    case 200:
      return { ErrorMessage: null, Result: true };
    case 400:
    case 500:
    case 409:
      return {
        ErrorMessage: ((await response.json()) as { error: string }).error,
        Result: false,
      };
    default:
      return { ErrorMessage: `[UNKNOW_ERROR]`, Result: false };
  }
}

export async function syncronizeExchangeAction(
  clusterId: number,
  exchangeId: number
): Promise<FrontResponse<boolean>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/exchange/${exchangeId}/syncronize`;
  console.log(
    `Sending request to syncronize exchange ${exchangeId} on cluster ${clusterId}`
  );

  let response = await fetch(createUserEndpoint, {
    method: "POST",
    cache: "no-store",
  });

  console.log(
    `Received response ${response.status} from ${createUserEndpoint}`
  );

  switch (response.status) {
    case 201:
    case 200:
      return { ErrorMessage: null, Result: true };
    case 400:
    case 500: {
      let contentBadRequest = (await response.json()) as { error: string };
      console.error(`Receiving error ${contentBadRequest.error}`);
      return { ErrorMessage: contentBadRequest.error, Result: false };
    }
    default:
      return { ErrorMessage: `[UNKNOW_ERROR]`, Result: false };
  }
}
