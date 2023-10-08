export type RabbitMqQueue = {
  ID: number;
  ClusterID: number;
  Type: string;
  Vhost: string;
  Arguments: Map<string, string>;
  Name: string;
  IsInCluster: boolean;
  IsInDatabase: boolean;
};
