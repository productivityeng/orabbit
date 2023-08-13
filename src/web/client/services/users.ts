"use server"

import { CreateRabbitmqUserSchema } from "@/schemas/user-schemas";
import { FrontResponse } from "./common/frontresponse";
import { RabbitMqCluster, RabbitMqUser } from "@/types";
import { z } from "zod";


export async function fetchUser(userId: number,clusterId: number) : Promise<FrontResponse<RabbitMqUser| null>> {
  const fetchUserEndpoint = `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}}/user/${userId}`
  let response = await fetch(fetchUserEndpoint,{
    method: "GET",
    cache:'no-store'
  })

  let contentResponse = await response.json() as RabbitMqUser;
  return {ErrorMessage: null,Result: contentResponse}
}

export async function createUser(request: z.infer<typeof CreateRabbitmqUserSchema>): Promise<FrontResponse<RabbitMqUser | null>> {
    const createUserEndpoint = `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/user`
    console.log("URL ",createUserEndpoint);
    let response = await fetch(createUserEndpoint, {
      body: JSON.stringify(request),
      method: "POST",
    });
    switch(response.status){
      case 201:{
        let contentResponse = await response.json() as RabbitMqUser;
        return {ErrorMessage: null,Result: contentResponse}
      }
      
      case 400:{
        let contentBadRequest = await response.json() as {error: string}
        return {ErrorMessage: contentBadRequest.error,Result: null}
      }

      case 406: {
        let contentInaceptable = await response.json() as {error: string,field: string}
        return {ErrorMessage: `field ${contentInaceptable.field} with error => ${contentInaceptable.error}`,Result: null}
      }

      case 500: {
        let contenctUnkow= await response.json()
        return {ErrorMessage: JSON.stringify(contenctUnkow),Result: null}
      }
      default:
        throw new Error("Erro desconhecido => "+ JSON.stringify(response))
    }   
  }