import { RabbitMqCluster } from "@/types";


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
    })
    let payloadResult = await result.json();
    let finalResult =  payloadResult as FetchAllClustersResult
    console.debug(finalResult)
    return finalResult;
}