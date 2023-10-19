"use server"
import { CreateRabbitMqClusterRequest } from "@/models/cluster";
import { RabbitMqCluster } from "@/types";
import { FrontResponse } from "./common/frontresponse";


export type FetchAllClustersResult =  {
    result: RabbitMqCluster[]
    pageNumber:1,
    pageSize: 100,
    totalItems: 3
}
export async function fetchAllClusters(){
    //todo: build a better method for retrieve all brokers
    let result = await fetch(`${process.env.PRIVATE_INVENTORY_ENDPOINT!}/broker?PageNumber=1&PageSize=100`,{
        method:'GET',
        cache:'no-store'
    })
    let payloadResult = await result.json();
    let finalResult =  payloadResult as FetchAllClustersResult
    return finalResult;
}

export async function createNewCluster(request: CreateRabbitMqClusterRequest): Promise<FrontResponse<RabbitMqCluster | null>> {
    let response = await fetch(`${process.env.PRIVATE_INVENTORY_ENDPOINT!}/broker`, {
      body: JSON.stringify(request),
      method: "POST",
    });
    switch(response.status){
      case 201:{
        let contentResponse = await response.json() as RabbitMqCluster;
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