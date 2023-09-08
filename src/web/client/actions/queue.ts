"use server";

import { RabbitMqQueue } from "@/types";
import { FrontResponse } from "./common/frontresponse";
import { CreateRabbitMqQeueueRequestSchema } from "@/schemas/queue-schemas";
import { z } from "zod";

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
