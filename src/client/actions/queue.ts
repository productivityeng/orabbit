"use server";

import { RabbitMqQueue } from "@/models/queues";
import { FrontResponse } from "./common/frontresponse";
import { CreateRabbitMqQeueueRequestSchema } from "@/schemas/queue-schemas";
import { boolean, z } from "zod";
import { log } from "console";

/**
 * Import a queue that exists in rabbitmq cluster and register it on ostern
 * @param clusterId
 * @param queueName
 */
export async function ImportQueueFromClusterAction(
  clusterId: number,
  VirtualHostName: string,
  queueName: string
) {
  const importQueueFromClusterEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/queue/import`;
  const body = JSON.stringify({
    QueueName: queueName,
    VirtualHostName: VirtualHostName,
  });

  console.info(
    `Sending request to ${importQueueFromClusterEndpoint} with body ${body}`
  );

  let result = await fetch(importQueueFromClusterEndpoint, {
    method: "POST",
    cache: "no-store",
    body,
  });
  console.info(`Receiving response ${result.status} `);
  switch (result.status) {
    case 201:
    case 200:
      return { ErrorMessage: null, Result: true };
    case 400:
    case 500: {
      const bodyResponse = (await result.json()) as { error: string };
      return { ErrorMessage: bodyResponse.error, Result: false };
    }
    default:
      return { ErrorMessage: `[UNKNOW_ERROR]`, Result: false };
  }
}

export async function fetchQeueusFromCluster(
  clusterId: number
): Promise<FrontResponse<RabbitMqQueue[] | null>> {
  const fetchQueueFromClusterEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/queue/queuesfromcluster`;

  let response = await fetch(fetchQueueFromClusterEndpoint, {
    method: "GET",
    cache: "no-store",
  });

  switch (response.status) {
    case 200:
      return {
        ErrorMessage: null,
        Result: (await response.json()) as RabbitMqQueue[],
      };
    default:
      return {
        ErrorMessage: ((await response.json()) as { error: string }).error,
        Result: null,
      };
  }
}

/**
 * Get specified detail about a RabbitMQ queue from a cluster
 * @param queueId
 * @param clusterId
 * @returns
 */
export async function fetchQueue(
  queueId: number,
  clusterId: number
): Promise<FrontResponse<RabbitMqQueue | null>> {
  const fetchUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/queue/${queueId}`;
  let response = await fetch(fetchUserEndpoint, {
    method: "GET",
    cache: "no-store",
  });

  switch (response.status) {
    case 404:
      return { ErrorMessage: "[QUEUE_NOT_FOUND]", Result: null };
  }

  let contentResponse = (await response.json()) as RabbitMqQueue;
  return { ErrorMessage: null, Result: contentResponse };
}

/**
 * Create a new Queue in specified RabbitMQ cluster and register it on ostern
 * @param clusterId C
 * @param request
 * @returns
 */
export async function createQueue(
  clusterId: number,
  request: z.infer<typeof CreateRabbitMqQeueueRequestSchema>
): Promise<FrontResponse<RabbitMqQueue | null>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/queue`;

  let response = await fetch(createUserEndpoint, {
    body: JSON.stringify(request),
    method: "POST",
    cache: "no-store",
  });

  switch (response.status) {
    case 201: {
      let contentResponse = (await response.json()) as RabbitMqQueue;
      return { ErrorMessage: null, Result: contentResponse };
    }

    case 400: {
      let contentBadRequest = (await response.json()) as { error: string };
      return { ErrorMessage: contentBadRequest.error, Result: null };
    }

    case 404: {
      let contentNotFoundRequest = (await response.json()) as { error: string };
      return { ErrorMessage: contentNotFoundRequest.error, Result: null };
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
      console.error(JSON.stringify(response));
      return { ErrorMessage: `[UNKNOW_ERROR]`, Result: null };
  }
}

export async function syncronizeQueueAction({
  ClusterId,
  QueueId,
}: {
  ClusterId: number;
  QueueId: number;
}): Promise<FrontResponse<boolean>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${ClusterId}/queue/syncronize`;
  console.log(
    `Sending request to sincronize queue ${QueueId} on cluster ${ClusterId}`
  );

  let response = await fetch(createUserEndpoint, {
    body: JSON.stringify({
      QueueId,
    }),
    method: "POST",
    cache: "no-store",
  });

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

/**
 * Removes a queue from a specified RabbitMQ Cluster based on ClusterID and QueueID.
 * @param param0
 * @returns
 */
export async function removeQueueFromClusterAction(
  ClusterId: number,
  QueueId: number
): Promise<FrontResponse<boolean>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${ClusterId}/queue/remove`;

  let response = await fetch(createUserEndpoint, {
    body: JSON.stringify({
      QueueId,
    }),
    method: "DELETE",
    cache: "no-store",
  });

  switch (response.status) {
    case 201:
    case 200:
      return { ErrorMessage: null, Result: true };
    case 400:
      return {
        ErrorMessage: await response.text(),
        Result: false,
      };
    default:
      return { ErrorMessage: `[UNKNOW_ERROR]`, Result: false };
  }
}
