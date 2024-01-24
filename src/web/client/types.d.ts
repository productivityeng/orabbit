export type RabbitMqCluster = {
  Id: number;
  CreatedAt: Date;
  UpdatedAt: Date;
  DeletedAt: Date | null;
  Name: string;
  Description: string;
  Host: string;
  Port: number;
  User: string;
  Password: string;
};
