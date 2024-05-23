import { render } from "@testing-library/react";
import { QueueTable } from "./queue-table";
import { RabbitMqQueueColumn } from "./columns";
import { GenerateFakeRabbitMqQueue } from "@/__mocks__/models-generator";



jest.mock('next/navigation', () => ({
    useRouter: () => ({
        query: { clusterId: '200' }
    }),
}));

jest.mock('next-intl',()=> ({
    useTranslations: () => jest.fn()
}))

describe('Queue table render state',() =>{
    it('should render queue table',() =>{
        const data = GenerateFakeRabbitMqQueue(10,200);
        const { getByText } = render(<QueueTable data={data} columns={RabbitMqQueueColumn} />);
    })
})