export type RabbitMqUser = {
  Id: number;
  BrokerId: number;
  Username: string;
  PasswordHash: string;
  IsInCluster: boolean;
  IsInDatabase: boolean;
};

export type ImportRabbitMqUser = {
  ClusterId: number;
  Username: string;
  Create: boolean;
};
