import { render, screen } from "@testing-library/react";
import { UserTable } from "./user-table";

import 'jest';
import assert from "assert";
import { GenerateRabbitMqList } from "@/__mocks__/models-generator";
import { RabbitMqUserTableColumnsDef } from "./columns";




jest.mock('next/navigation', () => ({
    useRouter: () => ({
        query: { clusterId: '200' }
    }),
  
}));

describe('UserTable render state', () => {
    let users = GenerateRabbitMqList(5,200);
    it('should render the component with correct users', async () => {
        
        render(<UserTable data={users} columns={RabbitMqUserTableColumnsDef} />);
        let userRows = await screen.findAllByTestId('user-table-row');
        expect(userRows).toHaveLength(users.length);
    });

    it('should render the component with all actions buttons disabled by default', async () => {
        const {findByTestId} = render(<UserTable data={users} columns={RabbitMqUserTableColumnsDef} />);
        let removeUserButton = await findByTestId('remove-user-button');
        let syncUserButton = await screen.findByTestId('syncronize-user-button');
        let importUserButton = await screen.findByTestId('import-user-button');
        let lockUserButton = await screen.findByTestId('lock-unlock-button');
        
        expect(removeUserButton).toHaveProperty('disabled', true);
        expect(syncUserButton).toHaveProperty('disabled', true);
        expect(importUserButton).toHaveProperty('disabled', true);
        expect(lockUserButton).toHaveProperty('disabled', true);
    });

});