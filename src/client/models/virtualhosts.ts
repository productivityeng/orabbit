import { LockerModel } from "@/actions/locker";

export interface RabbitMqVirtualHost {
  Id: number;
  Description: string;
  Name: string;
  ClusterId: number;
  DefaultQueueType: string;
  Tags: string[];
  IsInCluster: boolean;
  IsInDatabase: boolean;
  Lockers: LockerModel[];
}
