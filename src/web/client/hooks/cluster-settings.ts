import { create } from "zustand"


interface UseClusterSettingsStore {
    isDeleteModalOpen: boolean
    openDeleteModal: () => void
    closeDeleteModal: () => void
}


export const useClusterSettings = create<UseClusterSettingsStore>((set) => ({
    isDeleteModalOpen: false,
    closeDeleteModal: () => set({isDeleteModalOpen:false}),
    openDeleteModal: () => set({isDeleteModalOpen: true})
}))