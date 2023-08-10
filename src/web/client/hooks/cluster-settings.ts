import { DeleteCluster } from "@/services/cluster"
import { create } from "zustand"


interface UseClusterSettingsStore {
    isDeleteModalOpen: boolean
    openDeleteModal: () => void
    closeDeleteModal: () => void
    deleteCluster: (clusterId:number) => Promise<Boolean>
}


export const useClusterSettings = create<UseClusterSettingsStore>((set) => ({
    isDeleteModalOpen: false,
    closeDeleteModal: () => set({isDeleteModalOpen:false}),
    openDeleteModal: () => set({isDeleteModalOpen: true}),
    deleteCluster: DeleteCluster
}))