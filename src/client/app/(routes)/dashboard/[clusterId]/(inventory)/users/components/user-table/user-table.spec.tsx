import { act, fireEvent, render, screen, waitFor } from "@testing-library/react";
import { UserTable } from "./user-table";

import 'jest';
import assert from "assert";
import { GenerateRabbitMqList } from "@/__mocks__/models-generator";
import { RabbitMqUserTableColumnsDef } from "./columns";
import { CheckboxProps } from "@radix-ui/react-checkbox";
import { UserTableContext, UserTableContextProps } from "./user-table-context";

jest.mock('next/navigation', () => ({
    useRouter: () => ({
        query: { clusterId: '200' }
    }),
}));

jest.mock('next-intl',()=> ({
    useTranslations: () => jest.fn()
}))
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
     
        const {findByTestId} = render(
        
        <UserTableContext.Provider 
        value={{ onSyncronizeUser: SyncronizeUserActionMock,onRemoveUser: jest.fn() }}>
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

    it('should remove user when remove button is clicked', async () => {
            
            const removeUserFromClusterMock = jest.fn();
   
            const {findByTestId} = render(
            
            <UserTableContext.Provider 
            value={{ onRemoveUser: removeUserFromClusterMock,onSyncronizeUser: jest.fn() }}>
            <UserTable data={[{
                Id: 1,
                Username: 'test',
                PasswordHash: 'test',
                Lockers: [],
                ClusterId: 200,
                IsInCluster: true,
                IsInDatabase: true
            }]} columns={RabbitMqUserTableColumnsDef} /> </UserTableContext.Provider>);
            const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
            const removeUserButton = await findByTestId("remove-user-button");
    
            expect(removeUserButton).toHaveProperty('disabled', true);
            expect(userCheckbox).toBeDefined()
            act(() => {
                fireEvent.click(userCheckbox);
            });
            expect(removeUserButton).toHaveProperty('disabled', false);
            act(()=>{
                fireEvent.click(removeUserButton);
            });
            expect(removeUserFromClusterMock).toHaveBeenCalled();
    });

    it('should call import user when user is not tracked', async () => {
        const onImportUserMock = jest.fn();
        const {findByTestId} = render(
            <UserTableContext.Provider 
            value={{ onImportUser: onImportUserMock }}><UserTable data={[{
                Id: 1,
                Username: 'test',
                PasswordHash: 'test',
                Lockers: [],
                ClusterId: 200,
                IsInCluster: true,
                IsInDatabase: false
            }]} columns={RabbitMqUserTableColumnsDef} /></UserTableContext.Provider>
        );

        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const importUserButton = await findByTestId("import-user-button");


        expect(importUserButton).toHaveProperty('disabled', true);

        expect(userCheckbox).toBeDefined()
        act(() => {
            fireEvent.click(userCheckbox);
        });
        expect(importUserButton).toHaveProperty('disabled', false);
        act(()=>{
            fireEvent.click(importUserButton);
        });
        expect(onImportUserMock).toHaveBeenCalled();
    });

    it('should lock user when lock button is clicked', async () => {
        const onLockUserMock = jest.fn();
        const {findByTestId,getByTestId} = render(
            <UserTableContext.Provider 
            value={{ onLockUser: onLockUserMock }}><UserTable data={[{
                Id: 1,
                Username: 'test',
                PasswordHash: 'test',
                Lockers: [],
                ClusterId: 200,
                IsInCluster: true,
                IsInDatabase: true
            }]} columns={RabbitMqUserTableColumnsDef} /></UserTableContext.Provider>
        );

        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const lockUserButton = await findByTestId("lock-unlock-button");

        expect(lockUserButton).toHaveProperty('disabled', true);

        expect(userCheckbox).toBeDefined()
        act(() => {
            fireEvent.click(userCheckbox);
        });
        expect(lockUserButton).toHaveProperty('disabled', false);

        act(()=>{
            fireEvent.click(lockUserButton);
        });

        let lockItemDialog = await findByTestId("lock-item-dialog");
        expect(lockItemDialog).toBeDefined();

        const lockItemReasonTextarea = await findByTestId("lock-item-reason-textarea");
        act(()=>{
            fireEvent.change(lockItemReasonTextarea, { target: { value: 'reason with more than 10 characters' } });
        })

        const lockItemSubmitButton = await findByTestId("lock-item-submit-button");
        act(()=>{
            fireEvent.click(lockItemSubmitButton);
        });
        
        await waitFor(() => {
            expect(onLockUserMock).toHaveBeenCalled();
        });
     });

    it('should unlock user when unlock button is clicked', async () => { 
        const onUnlockUserMock = jest.fn();
        const {findByTestId} = render(
            <UserTableContext.Provider 
            value={{ onUnlockUser: onUnlockUserMock }}><UserTable data={[{
                Id: 1,
                Username: 'test',
                PasswordHash: 'test',
                Lockers: [{
                    Id: 1,
                    Reason: 'reason',
                    UserResponsibleEmail: 'user@email.com',
                    CreatedAt: new Date(),
                    Enabled: true,
                    UpdatedAt: new Date(),
                    
                }],
                ClusterId: 200,
                IsInCluster: true,
                IsInDatabase: true
            }]} columns={RabbitMqUserTableColumnsDef} /></UserTableContext.Provider>
        );

        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const unlockIconButton = await findByTestId("unlock-icon-button");

        act(() => {
            fireEvent.click(userCheckbox);
            fireEvent.click(unlockIconButton);
        });

        await findByTestId("unlock-dialog");
        const unlockUserButton = await findByTestId("unlock-action-button");

        act(()=>{
            fireEvent.click(unlockUserButton);
        });

        await waitFor(() => {
            expect(onUnlockUserMock).toHaveBeenCalledTimes(1);
        });
    });

    it('should not show any action if user is locked', async () => {
        const {findByTestId} = render(
            <UserTable data={[{
                Id: 1,
                Username: 'test',
                PasswordHash: 'test',
                Lockers: [{
                    Id: 1,
                    Reason: 'reason',
                    UserResponsibleEmail: 'user',
                    CreatedAt: new Date(),
                    Enabled: true,
                    UpdatedAt: new Date(),
                    
                }],
                ClusterId: 200,
                IsInCluster: true,
                IsInDatabase: true
            }]} columns={RabbitMqUserTableColumnsDef} />
        );

        const userCheckbox = await findByTestId(`user-table-checkbox-${1}`);
        const lockUserButton = await findByTestId("lock-unlock-button");
        const removeUserButton = await findByTestId("remove-user-button");
        const syncUserButton = await findByTestId("syncronize-user-button");
        const importUserButton = await findByTestId("import-user-button");

        expect(lockUserButton).toHaveProperty('disabled', true);
        expect(removeUserButton).toHaveProperty('disabled', true);
        expect(syncUserButton).toHaveProperty('disabled', true);
        expect(importUserButton).toHaveProperty('disabled', true);

        act(() => {
            fireEvent.click(userCheckbox);
        });

        expect(lockUserButton).toHaveProperty('disabled', true);
        expect(removeUserButton).toHaveProperty('disabled', true);
        expect(syncUserButton).toHaveProperty('disabled', true);
        expect(importUserButton).toHaveProperty('disabled', true);
    });

});