

export interface FrontResponse<T> {
    Result: T
    ErrorMessage: string | null
}

export interface PaginatedResponse<T>{
    result: T[]
    pageNumber: number,
    pageSize: number,
    totalItems: number
}