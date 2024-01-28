"use server";

import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import { FrontResponse } from "./common/frontresponse";

export async function fetchVirtualHosts(clusterId: number) {
  let result = await fetch(
    `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/virtualhost`,
    {
      method: "GET",
      cache: "no-store",
    }
  );
  let payloadResult = await result.json();
  let finalResult = payloadResult as RabbitMqVirtualHost[];
  return finalResult;
}

export async function removeVirtualHostAction(
  clusterId: number,
  virtualHostId: number
): Promise<FrontResponse<RabbitMqVirtualHost | undefined>> {
  let result = await fetch(
    `${process.env
      .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/virtualhost/${virtualHostId}`,
    {
      method: "DELETE",
      cache: "no-store",
    }
  );

  if (result.status !== 200) {
    return {
      ErrorMessage: ((await result.json()) as { error: string }).error,
      Result: undefined,
    };
  }

  let payloadResult = await result.json();
  let finalResult = payloadResult as RabbitMqVirtualHost;

  return {
    ErrorMessage: null,
    Result: finalResult,
  };
}

export async function syncronizeVirtualHostAction(
  clusterId: number,
  virtualHostId: number
): Promise<FrontResponse<RabbitMqVirtualHost | undefined>> {
  let result = await fetch(
    `${process.env
      .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/virtualhost/${virtualHostId}/syncronize`,
    {
      method: "POST",
      cache: "no-store",
    }
  );

  switch (result.status) {
    case 200:
    case 201: {
      let payloadResult = await result.json();
      let finalResult = payloadResult as RabbitMqVirtualHost;

      return {
        ErrorMessage: null,
        Result: finalResult,
      };
    }
    default: {
      return {
        ErrorMessage: ((await result.json()) as { error: string }).error,
        Result: undefined,
      };
    }
  }
}
