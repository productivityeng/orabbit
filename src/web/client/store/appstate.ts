import { RabbitMqCluster } from "@/types";
import { create } from "zustand";


interface AppState {
    AvailableClusters: RabbitMqCluster[] | undefined
    SelectedClusterId: number | undefined
    SetAvailableClusters: (clusters: RabbitMqCluster[]) =>  void
    SetSelectedClusterId: (clusterid: number| undefined) => void
    GetSelectedCluster: () => RabbitMqCluster | undefined,
}

export const useAppState = create<AppState>((set,get)=> ({
    AvailableClusters: undefined,
    SelectedClusterId: undefined,
    GetSelectedCluster: () => {
        return get().AvailableClusters?.find(x => x.Id == get().SelectedClusterId)
    },
    SetSelectedClusterId : (clusterId:number| undefined) =>{
        set({SelectedClusterId: clusterId})
    },
    SetAvailableClusters: (clusters: RabbitMqCluster[]) => {
        set({AvailableClusters:clusters})
    }   
}))


