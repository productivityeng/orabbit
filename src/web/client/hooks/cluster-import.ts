import { CreateRabbitMqClusterRequestSchema } from "@/schemas/cluster-schemas";
import { createNewCluster } from "@/actions/cluster";
import { FrontResponse } from "@/actions/common/frontresponse";
import { RabbitMqCluster } from "@/types";
import { z } from "zod";
import { create } from "zustand";


interface ImportClusterState {
   isModalOpen: boolean
   openModal: () => void
   closeModal: () => void,
   importCluster: (values: z.infer<typeof CreateRabbitMqClusterRequestSchema>) => Promise<FrontResponse<RabbitMqCluster | null>>
}

export const useImportCluster = create<ImportClusterState>((set,get)=> ({
   isModalOpen: false,
   openModal: () => set({isModalOpen:true}),
   closeModal: () => set({isModalOpen:false}),
   importCluster: createNewCluster
}))


