
export type RabbitMqCluster = {
    Id: number
    createdAt: Date
    updatedAt:Date
    deletedAt: Date | null
    name: string
    description: string
    host: string
    port: number
    user: string
    password: string
}

export type RabbitMqUser = {
    Id: number
    BrokerId: number
    Username: string
    PasswordHash: string
}
