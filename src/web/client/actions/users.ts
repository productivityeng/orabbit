"use server";

import { CreateRabbitmqUserSchema } from "@/schemas/user-schemas";
import { FrontResponse, PaginatedResponse } from "./common/frontresponse";
import { ImportRabbitMqUser, RabbitMqCluster, RabbitMqUser } from "@/types";
import { z } from "zod";

/**
 *
 * @param clusterId id of a broker where from the user will be searched
 * @param page number of a page to be retrieved from server
 * @param pagesize length of each page
 * @returns
 */
export async function fetchRegisteredUsers(
  clusterId: number,
  page: number = 1,
  pagesize: number = 10
) {
  let result = await fetch(
    `${process.env
      .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/user?PageNumber=${page}&PageSize=${pagesize}`,
    {
      method: "GET",
      cache: "no-store",
    }
  );
  let payloadResult = await result.json();
  let finalResult = payloadResult as PaginatedResponse<RabbitMqUser>;
  return finalResult;
}

export async function importUserFromCluster(
  request: ImportRabbitMqUser
): Promise<RabbitMqUser> {
  let result = await fetch(
    `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/${request.ClusterId}/user`,
    {
      method: "POST",
      cache: "no-store",
      body: JSON.stringify(request),
    }
  );

  let payloadResult = await result.json();
  return payloadResult;
}

/**
 *
 * @param clusterId id of a broker where from the user will be searched
 * @returns
 */
export async function fetchUsersFromCluster(
  clusterId: number
): Promise<RabbitMqUser[]> {
  let result = await fetch(
    `${process.env
      .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/user/usersfromcluster`,
    {
      method: "GET",
      cache: "no-store",
    }
  );
  let payloadResult = await result.json();
  return payloadResult;
}

export async function fetchUser(
  userId: number,
  clusterId: number
): Promise<FrontResponse<RabbitMqUser | null>> {
  const fetchUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/user/${userId}`;
  let response = await fetch(fetchUserEndpoint, {
    method: "GET",
    cache: "no-store",
  });

  let contentResponse = (await response.json()) as RabbitMqUser;
  return { ErrorMessage: null, Result: contentResponse };
}

export async function createUser(
  clusterId: number,
  request: z.infer<typeof CreateRabbitmqUserSchema>
): Promise<FrontResponse<RabbitMqUser | null>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/user`;
  console.log("URL ", createUserEndpoint);
  let response = await fetch(createUserEndpoint, {
    body: JSON.stringify(request),
    method: "POST",
  });
  switch (response.status) {
    case 201: {
      let contentResponse = (await response.json()) as RabbitMqUser;
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

/**
 * Remove o usuario da base do ostern e tambem do rabbitmq.
 * @param clusterId Identificação do Cluster do usuario
 * @param userId Identificação Global do usuário
 * @returns
 */
export async function deleteUserFromRabbit(
  clusterId: number,
  userId: number
): Promise<FrontResponse<string | null>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/user/rabbitmq/${userId}`;

  let response = await fetch(createUserEndpoint, {
    method: "DELETE",
  });

  switch (response.status) {
    case 204: {
      return { ErrorMessage: null, Result: "Deleted" };
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

/**
 * Remove o usuario da base do ostern, nao do rabbitmq.
 * @param clusterId Identificação do Cluster do usuario
 * @param userId Identificação Global do usuário
 * @returns
 */
export async function deleteUserFromTracking(
  clusterId: number,
  userId: number
): Promise<FrontResponse<string | null>> {
  const createUserEndpoint = `${process.env
    .PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/user/${userId}`;

  let response = await fetch(createUserEndpoint, {
    method: "DELETE",
  });

  switch (response.status) {
    case 204: {
      return { ErrorMessage: null, Result: "Deleted" };
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
