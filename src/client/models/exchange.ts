import { LockerModel } from "@/actions/locker";

export type RabbitMqExchange = {
  Id: number;
  Name: string;
  ClusterId: number;
  Internal: boolean;
  Durable: boolean;
  Arguments: Map<string, string>;
  Lockers: LockerModel[];
  VHost: string;
  Type: string;
  IsInCluster: boolean;
  IsInDatabase: boolean;
};
