export type RabbitMqCluster = {
  ID: number;
  CreatedAt: Date;
  updatedAt: Date;
  deletedAt: Date | null;
  name: string;
  description: string;
  host: string;
  port: number;
  user: string;
  password: string;
};

