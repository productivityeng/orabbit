import { RabbitMqQueue } from "@/models/queues";
import { RabbitMqUser } from "@/models/users";
import { RabbitMqVirtualHost } from "@/models/virtualhosts";
import { faker } from "@faker-js/faker";

export function GenerateRabbitMqList(count: number, clusterNumber: 200) {
  const list: RabbitMqUser[] = [];

  for (let i = 0; i < count; i++) {
    const rabbitMQ: RabbitMqUser = {
      ClusterId: clusterNumber,
      Id: faker.number.int(),
      IsInCluster: faker.datatype.boolean(),
      IsInDatabase: faker.datatype.boolean(),
      PasswordHash: faker.internet.password(),
      Username: faker.internet.userName(),
      Lockers: [],
    };

    list.push(rabbitMQ);
  }

  return list;
}

export function GenerateFakeRabbitMqQueue(
  count: number,
  clusterNumber: number
) {
  const list: RabbitMqQueue[] = [];

  for (let i = 0; i < count; i++) {
    const queue: RabbitMqQueue = {
      Arguments: new Map<string, string>(),
      ClusterId: clusterNumber,
      ID: faker.number.int(),
      Name: faker.internet.userName(),
      Durable: faker.datatype.boolean(),
      IsInCluster: faker.datatype.boolean(),
      IsInDatabase: faker.datatype.boolean(),
      Lockers: [],
      Type: "classic",
      VHost: faker.internet.domainName(),
    };

    list.push(queue);
  }

  return list;
}

export function GenerateVirtualHostMq(count: number, clusterNumber: number) {
  const list: RabbitMqVirtualHost[] = [];
  for (let i = 0; i < count; i++) {
    const virtualHost: RabbitMqVirtualHost = {
      ClusterId: clusterNumber,
      Id: faker.number.int(),
      Name: faker.internet.domainName(),
      IsInCluster: faker.datatype.boolean(),
      IsInDatabase: faker.datatype.boolean(),
      Lockers: [],
      DefaultQueueType: "classic",
      Description: faker.lorem.sentence(),
      Tags: [],
    };
    list.push(virtualHost);
  }
  return list;
}
