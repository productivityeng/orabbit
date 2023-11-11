import { PaginatedResponse } from "@/actions/common/frontresponse";
import { VirtualHosts } from "../models/virtualhosts";

export async function fetchVirtualHosts(clusterId: number) {
  let result = await fetch(
    `${process.env.PRIVATE_INVENTORY_ENDPOINT!}/${clusterId}/virtualhost`,
    {
      method: "GET",
      cache: "no-store",
    }
  );
  let payloadResult = await result.json();
  let finalResult = payloadResult as VirtualHosts[];
  return finalResult;
}
