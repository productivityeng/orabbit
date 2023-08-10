import { create } from "zustand";


interface ImportClusterState {
   isModalOpen: boolean
   openModal: () => void
   closeModal: () => void
}

export const useImportCluster = create<ImportClusterState>((set,get)=> ({
   isModalOpen: false,
   openModal: () => set({isModalOpen:true}),
   closeModal: () => set({isModalOpen:false})
}))


