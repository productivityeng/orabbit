import { LockerModel } from "@/actions/locker";

export type RabbitMqUser = {
  Id: number;
  ClusterId: number;
  Username: string;
  PasswordHash: string;
  IsInCluster: boolean;
  IsInDatabase: boolean;
  Lockers: LockerModel[];
};

export type ImportRabbitMqUser = {
  ClusterId: number;
  Username: string;
  Create: boolean;
};
