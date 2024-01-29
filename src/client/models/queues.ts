import { LockerModel } from "@/actions/locker";

export type RabbitMqQueue = {
  ID: number;
  ClusterId: number;
  Type: string;
  VHost: string;
  Arguments: Map<string, string>;
  Name: string;
  IsInCluster: boolean;
  IsInDatabase: boolean;
  Durable: boolean;
  Lockers: LockerModel[];
};
