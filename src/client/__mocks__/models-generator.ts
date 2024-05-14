import { RabbitMqUser } from "@/models/users";
import {faker} from '@faker-js/faker';

export function GenerateRabbitMqList(count: number,clusterNumber:200) {
    const list:RabbitMqUser[] = [];

    for (let i = 0; i < count; i++) {
        const rabbitMQ:RabbitMqUser = {
            ClusterId: clusterNumber,
            Id: faker.number.int(),
            IsInCluster: faker.datatype.boolean(),
            IsInDatabase: faker.datatype.boolean(),
            PasswordHash: faker.internet.password(),
            Username: faker.internet.userName(),
            Lockers:[]
        };

        list.push(rabbitMQ);
    }

    return list;
}