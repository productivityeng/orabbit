import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { UserTable } from "./user-table";

import 'jest';
import assert from "assert";
import { GenerateRabbitMqList } from "@/__mocks__/models-generator";
import { RabbitMqUserTableColumnsDef } from "./columns";
import { CheckboxProps } from "@radix-ui/react-checkbox";
import { act } from "react-dom/test-utils";
import { UserTableContext, UserTableContextProps } from "./user-table-context";




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

    it('should enabled loack and remove button when a unlocked user is selected', async () => {
        const {findByTestId} = render(<UserTable data={[{
            Id: 1,
            Username: 'test',
            PasswordHash: 'test',
            Lockers: [],
            ClusterId: 200,
            IsInCluster: true,
            IsInDatabase: true
        }]} columns={RabbitMqUserTableColumnsDef} />);

        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const lockUserButton = await findByTestId("lock-unlock-button");
        const removeUserButton = await findByTestId("remove-user-button");

        expect(lockUserButton).toHaveProperty('disabled', true);
        expect(removeUserButton).toHaveProperty('disabled', true);

        expect(userCheckbox).toBeDefined()
        act(() => {
            userCheckbox.click();
        });
        
        expect(lockUserButton).toHaveProperty('disabled', false);
        expect(removeUserButton).toHaveProperty('disabled', false);

    });

    it('should cal remove user action when remove button is clicked', async () => {

        const removeUserFromClusterMock = jest.fn()
        jest.mock('../../../../../../../../actions/users', () => ({
            ...jest.requireActual('../../../../../../../../actions/users'),
            removeUserFromCluster: removeUserFromClusterMock,
          
        }));
        const {findByTestId} = render(<UserTable data={[{
            Id: 1,
            Username: 'test',
            PasswordHash: 'test',
            Lockers: [],
            ClusterId: 200,
            IsInCluster: true,
            IsInDatabase: true
        }]} columns={RabbitMqUserTableColumnsDef} />);
        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const removeUserButton = await findByTestId("remove-user-button");

        expect(userCheckbox).toBeDefined()
        act(() => {
            userCheckbox.click();
            removeUserButton.click();
            
        });
    });

    it('should sync user when sync button is clicked', async () => {
        
        const SyncronizeUserActionMock = jest.fn();
        jest.mock('@/actions/users', () => ({
            ...jest.requireActual('@/actions/users'),
            SyncronizeUserAction: SyncronizeUserActionMock,
        }));
     
        const {findByTestId} = render(
        
        <UserTableContext.Provider value={{ onSyncronizeUserClick: SyncronizeUserActionMock }}>
        <UserTable data={[{
            Id: 1,
            Username: 'test',
            PasswordHash: 'test',
            Lockers: [],
            ClusterId: 200,
            IsInCluster: false,
            IsInDatabase: true
        }]} columns={RabbitMqUserTableColumnsDef} /> </UserTableContext.Provider>);
        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const syncUserButton = await findByTestId("syncronize-user-button");

        expect(syncUserButton).toHaveProperty('disabled', true);
        expect(userCheckbox).toBeDefined()
        act(() => {
            fireEvent.click(userCheckbox);
        });
        expect(syncUserButton).toHaveProperty('disabled', false);
        act(()=>{
            fireEvent.click(syncUserButton);
        });
        expect(SyncronizeUserActionMock).toHaveBeenCalled();
       
    });

});