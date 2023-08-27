
export type RabbitMqCluster = {
    ID: number
    CreatedAt: Date
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
    IsRegistered: boolean
}

export type ImportRabbitMqUser = {
    ClusterId: number
    Username: string
    Create: boolean = false
}